package main

import (
	"fmt"

	"github.com/ervinismu/purplestore/internal/app/controller"
	"github.com/ervinismu/purplestore/internal/app/repository"
	"github.com/ervinismu/purplestore/internal/app/service"
	"github.com/ervinismu/purplestore/internal/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var cfg config.Config
var dbConn *sqlx.DB
var err error

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
}

func main() {
	r := gin.Default()

	// init repo
  categoryRepository := repository.NewCategoryRepository(dbConn)

	// init service
  categoryService := service.NewCategorySerivce(categoryRepository)

  // init controller
  categoryCotroller := controller.NewCategoryController(categoryService)

	// categories routes
  r.GET("/categories", categoryCotroller.GetList)
  r.POST("/categories", categoryCotroller.Create)

	// run server
	appPort := fmt.Sprintf(":%s", cfg.AppPort)
	r.Run(appPort)
}
