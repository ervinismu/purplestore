//go:build wireinject
// +build wireinject

package main

import (
	"github.com/ervinismu/purplestore/controllers"
	"github.com/ervinismu/purplestore/repository"
	"github.com/ervinismu/purplestore/service"
	"github.com/google/wire"
	"gorm.io/gorm"
)


func InitProductController(db *gorm.DB) controllers.ProductController {
	wire.Build(
		controllers.NewProductController,
		service.NewProductService,
		repository.NewProductRepository,
	)

	return controllers.ProductController{}
}

func InitUserController(db *gorm.DB) controllers.UserController {
	wire.Build(
		controllers.NewUserController,
		service.NewUserService,
		repository.NewUserRepository,
	)

	return controllers.UserController{}
}
