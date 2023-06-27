package controller

import (
	"net/http"
	"strconv"

	"github.com/ervinismu/purplestore/internal/app/schema"
	"github.com/ervinismu/purplestore/internal/pkg/handler"
	"github.com/ervinismu/purplestore/internal/pkg/reason"
	"github.com/gin-gonic/gin"
)

type SessionService interface {
	Login(req *schema.LoginRequest) (schema.LoginResp, error)
	Logout(UserID int) error
	Refresh(req *schema.RefreshTokenRequest) (schema.RefreshTokenResp, error)
}

type RefreshTokenVerifier interface {
	VerifyRefreshToken(tokenString string) (string, error)
}

type SessionController struct {
	service    SessionService
	tokenMaker RefreshTokenVerifier
}

func NewSessionController(service SessionService, tokenMaker RefreshTokenVerifier) *SessionController {
	return &SessionController{
		service:    service,
		tokenMaker: tokenMaker,
	}
}

func (ctrl *SessionController) Login(ctx *gin.Context) {
	req := &schema.LoginRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	resp, err := ctrl.service.Login(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "", resp)
}

// refresh
func (ctrl *SessionController) Refresh(ctx *gin.Context) {
	refreshToken := ctx.GetHeader("refresh_token")
	if refreshToken == "" {
		handler.ResponseError(ctx, http.StatusUnauthorized, reason.SessionFailedRefresh)
		return
	}

	sub, err := ctrl.tokenMaker.VerifyRefreshToken(refreshToken)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnauthorized, reason.SessionFailedRefresh)
		return
	}

	intSub, _ := strconv.Atoi(sub)
	req := &schema.RefreshTokenRequest{}
	req.RefreshToken = refreshToken
	req.UserID = intSub

	resp, err := ctrl.service.Refresh(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnauthorized, reason.SessionFailedRefresh)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "", resp)
}

// logout
func (ctrl *SessionController) Logout(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.GetString("user_id"))
	err := ctrl.service.Logout(userID)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, reason.SessionSuccessLogout, nil)
}
