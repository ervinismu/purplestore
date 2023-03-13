package router

import (
	"github.com/ervinismu/purplestore/controllers"
	"github.com/gin-gonic/gin"
)

type ProductApiRouter struct {
	productController controllers.ProductController
}

func NewProductApiRouter(productController controllers.ProductController) ProductApiRouter {
	return ProductApiRouter{productController: productController}
}

func (p ProductApiRouter) RegisterProductApiRouter(r *gin.RouterGroup) {
	r.GET("/products", p.productController.ListProducts)
	r.GET("/products/:id", p.productController.ShowProduct)
	r.POST("/products", p.productController.CreateProduct)
	r.DELETE("/products/:id", p.productController.DeleteProduct)
	r.PATCH("/products/:id", p.productController.UpdateProduct)
}
