package schema

type CreateUserReq struct {
	Email  string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type UpdateUserReq struct {
	Username string `json:"username"`
}
