package controllers

import (
	"github.com/ervinismu/purplestore/handler"
	"github.com/ervinismu/purplestore/schema"
	"github.com/ervinismu/purplestore/service"
	"github.com/gin-gonic/gin"
)

type productController struct {
	service service.ProductService
}

func NewProductController(ps service.ProductService) productController {
	return productController{ service: ps }
}

func (pc productController) ListProducts(ctx *gin.Context) {
	products, err := pc.service.ListAllProduct()
	if err != nil {
		handler.HandlerErrorResponse(ctx, "Cannot get list products")
		return
	}

	handler.HandlerSuccessResponse(ctx, "Success get products", products )
}

func (pc productController) ShowProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := pc.service.ShowDetailProduct(id)
	if err != nil {
		handler.HandlerErrorResponse(ctx, "Record not found")
		return
	}

	handler.HandlerSuccessResponse(ctx, "Success get detail products", product )
}

func (pc productController) CreateProduct(ctx *gin.Context) {
	var input schema.CreateProductReq
	if err := ctx.ShouldBindJSON(&input); err != nil {
		handler.HandlerErrorResponse(ctx, err.Error())
		return
	}
	product, err := pc.service.CreateProduct(input)
	if err != nil {
		handler.HandlerErrorResponse(ctx, err.Error())
		return
	}

	handler.HandlerSuccessResponse(ctx, "Success create product", product )
}

func (pc productController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := pc.service.DeleteProduct(id)
	if err != nil {
		handler.HandlerErrorResponse(ctx, "Record not found")
		return
	}

	handler.HandlerSuccessResponse(ctx, "Success delete product", product )
}

func (pc productController) UpdateProduct(ctx *gin.Context) {
	var input schema.UpdateProductReq
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&input); err != nil {
		handler.HandlerErrorResponse(ctx, err.Error())
		return
	}

	product, err := pc.service.UpdateProduct(id, input)
	if err != nil {
		handler.HandlerErrorResponse(ctx, "Record not found")
		return
	}

	handler.HandlerSuccessResponse(ctx, "Success update product", product)
}
