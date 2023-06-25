package controller

import (
	"net/http"
	"strconv"

	"github.com/ervinismu/purplestore/internal/app/schema"
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
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	resp, err := ctrl.service.Login(req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

// refresh
func (ctrl *SessionController) Refresh(ctx *gin.Context) {
	refreshToken := ctx.GetHeader("refresh_token")
	if refreshToken == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "cannot refresh token"})
		return
	}

	sub, err := ctrl.tokenMaker.VerifyRefreshToken(refreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "failed to refresh token"})
		return
	}

	intSub, _ := strconv.Atoi(sub)
	req := &schema.RefreshTokenRequest{}
	req.RefreshToken = refreshToken
	req.UserID = intSub

	resp, err := ctrl.service.Refresh(req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "failed to refresh token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

// logout
func (ctrl *SessionController) Logout(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.GetString("user_id"))
	err := ctrl.service.Logout(userID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success logout"})
}
