package controller

import (
	"net/http"

	"github.com/ervinismu/purplestore/internal/app/schema"
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
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	err = ctrl.service.Register(req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success register"})
}
