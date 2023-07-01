package controller

import (
	"net/http"

	"github.com/ervinismu/purplestore/internal/app/schema"
	"github.com/ervinismu/purplestore/internal/pkg/handler"
	"github.com/ervinismu/purplestore/internal/pkg/reason"
	"github.com/gin-gonic/gin"
)

type Registerer interface {
	Register(req *schema.RegisterRequest) error
}

type RegistrationController struct {
	service Registerer
}

func NewRegistrationController(service Registerer) *RegistrationController {
	return &RegistrationController{service: service}
}

func (ctrl *RegistrationController) Register(ctx *gin.Context) {
	req := &schema.RegisterRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	err = ctrl.service.Register(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, reason.RegisterSuccess, nil)
}
