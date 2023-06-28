package controller

import (
	"net/http"
	"strconv"

	"github.com/ervinismu/purplestore/internal/app/schema"
	"github.com/ervinismu/purplestore/internal/pkg/handler"
	"github.com/ervinismu/purplestore/internal/pkg/reason"
	"github.com/gin-gonic/gin"
)

type CategoryService interface {
	GetList() ([]schema.CategoryListResponse, error)
	Create(req schema.CategoryCreateRequest) error
	Detail(req schema.CategoryDetailRequest) (schema.CategoryDetailResponse, error)
	DeleteByID(req schema.CategoryDeleteRequest) error
	Update(req schema.CategoryUpdateRequest) error
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
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "", response)
}

func (ctrl *CategoryController) Create(ctx *gin.Context) {
	var req schema.CategoryCreateRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	err = ctrl.service.Create(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, reason.CategorySuccessCreate, nil)
}

func (ctrl *CategoryController) Detail(ctx *gin.Context) {
	categoryIDstr := ctx.Param("id")
	categoryID, err := strconv.Atoi(categoryIDstr)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.CategoryFailedGetDetail)
		return
	}

	req := schema.CategoryDetailRequest{ID: categoryID}
	response, err := ctrl.service.Detail(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "", response)
}

func (cc *CategoryController) Delete(ctx *gin.Context) {
	categoryIDstr := ctx.Param("id")
	categoryID, err := strconv.Atoi(categoryIDstr)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.CategoryFailedDelete)
		return
	}

	req := schema.CategoryDeleteRequest{ID: categoryID}
	err = cc.service.DeleteByID(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, reason.CategorySuccessDelete, nil)
}

func (cc *CategoryController) Update(ctx *gin.Context) {
	categoryIDstr := ctx.Param("id")
	categoryID, err := strconv.Atoi(categoryIDstr)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.CategoryFailedUpdate)
		return
	}

	req := schema.CategoryUpdateRequest{}
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.CategoryFailedUpdate)
		return
	}

	req.ID = categoryID

	err = cc.service.Update(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, reason.CategorySuccessUpdate, nil)
}
