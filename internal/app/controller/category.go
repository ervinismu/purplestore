package controller

import (
	"net/http"
	"strconv"

	"github.com/ervinismu/purplestore/internal/app/schema"
	"github.com/gin-gonic/gin"
)

type CategoryService interface {
	GetList() ([]schema.CategoryListResponse, error)
	Create(req schema.CategoryCreateRequest) error
	Detail(req schema.CategoryDetailRequest) (schema.CategoryDetailResponse, error)
}

type CategoryController struct {
	service CategoryService
}

func NewCategoryController(service CategoryService) *CategoryController {
	return &CategoryController{service: service}
}

func (ctrl *CategoryController) GetList(ctx *gin.Context) {
	response, err := ctrl.service.GetList()
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": response})
}

func (ctrl *CategoryController) Create(ctx *gin.Context) {
	var req schema.CategoryCreateRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	err = ctrl.service.Create(req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "failed create category"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "success create category"})
}

func (ctrl *CategoryController) Detail(ctx *gin.Context) {
	categoryIDstr := ctx.Param("id")
	categoryID, err := strconv.Atoi(categoryIDstr)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "failed get detail category"})
		return
	}

	req := schema.CategoryDetailRequest{ID: categoryID}
	response, err := ctrl.service.Detail(req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "failed get detail category"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": response})
}
