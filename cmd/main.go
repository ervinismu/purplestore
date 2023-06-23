package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/ervinismu/purplestore/internal/app/controller"
	"github.com/ervinismu/purplestore/internal/app/repository"
	"github.com/ervinismu/purplestore/internal/app/service"
	"github.com/ervinismu/purplestore/internal/pkg/config"
	"github.com/ervinismu/purplestore/internal/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var cfg config.Config
var dbConn *sqlx.DB
var err error
var enforcer *casbin.Enforcer

func init() {
	// load configuration based on app.env
	cfg, err = config.LoadConfig()
	if err != nil {
		panic("failed to load config")
	}

	// Create database connection
	dbConn, err = sqlx.Open(cfg.DatabaseDriver, cfg.DatabaseURL)
	if err != nil {
		errMsg := fmt.Errorf("err database connect: %w", err)
		panic(errMsg)
	}

	err = dbConn.Ping()
	if err != nil {
		errMsg := fmt.Errorf("err database ping: %w", err)
		panic(errMsg)
	}

	// cabin enforcer
	enforcer, err = casbin.NewEnforcer("config/rbac_model.conf", "config/rbac_policy.csv")
	if err != nil {
		errMsg := fmt.Errorf("err casbin enforcer: %w", err)
		panic(errMsg)
	}

	// setup logrus logging
	logLevel, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	r := gin.New()
	r.Use(middleware.LogMiddleware())

	// init repo
	categoryRepository := repository.NewCategoryRepository(dbConn)

	// init service
	categoryService := service.NewCategorySerivce(categoryRepository)

	// init controller
	categoryCotroller := controller.NewCategoryController(categoryService)

	// categories routes
	v1Routes := r.Group("api/v1")
	{
		v1Routes.GET("/categories", categoryCotroller.GetList)
		v1Routes.GET("/categories/:id", categoryCotroller.Detail)
		v1Routes.POST("/categories",
			middleware.AuthorizationMiddleware("bob", "categories", "write", enforcer),
			categoryCotroller.Create,
		)
		v1Routes.DELETE("/categories/:id",
			middleware.AuthorizationMiddleware("bob", "categories", "delete", enforcer),
			categoryCotroller.Delete,
		)
		v1Routes.PATCH("/categories/:id",
			middleware.AuthorizationMiddleware("bob", "categories", "patch", enforcer),
			categoryCotroller.Update,
		)
	}

	// run server
	appPort := fmt.Sprintf(":%s", cfg.AppPort)
	err := r.Run(appPort)
	if err != nil {
		log.Panic("cannot start app")
	}
}
