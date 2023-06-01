package controller

import (
	"net/http"

	"github.com/ervinismu/purplestore/internal/app/schema"
	"github.com/ervinismu/purplestore/internal/pkg/reason"
	"github.com/gin-gonic/gin"
)

type CategoryService interface {
	GetList() ([]schema.CategoryGetListResponse, error)
	Create(req schema.CategoryCreateRequest) error
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
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": response})
}

func (ctrl *CategoryController) Create(ctx *gin.Context) {
	var req schema.CategoryCreateRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": reason.CategoryCreateFailed})
		return
	}

	err = ctrl.service.Create(req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
