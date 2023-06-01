package main

import (
	"fmt"
	"net/http"

	"github.com/ervinismu/purplestore/internal/app/controller"
	"github.com/ervinismu/purplestore/internal/app/repository"
	"github.com/ervinismu/purplestore/internal/app/service"
	"github.com/ervinismu/purplestore/internal/pkg/config"
	"github.com/ervinismu/purplestore/internal/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var dbConn *sqlx.DB
var cfg config.Config
var err error

func init() {
	cfg, err = config.LoadConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg.DatabasDriver)
	fmt.Println(cfg.DatabaseURL)

	dbConn, err = db.ConnectDB(cfg.DatabasDriver, cfg.DatabaseURL)
	if err != nil {
		panic(err)
	}
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	// setup repository
	categoryRepo := repository.NewCategoryRepository(dbConn)
	// setup service
	categoryService := service.NewCategoryService(categoryRepo)
	// setup controller
	categoryController := controller.NewCategoryController(categoryService)

	// categories endpoint
	router.GET("/categories", categoryController.GetList)
	router.POST("/categories", categoryController.Create)

	router.Run()
}
