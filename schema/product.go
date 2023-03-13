package schema

type CreateProductReq struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateProductReq struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}
