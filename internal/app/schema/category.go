package schema

type CategoryGetListResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryCreateRequest struct {
	Name        string `binding:"required"`
	Description string `binding:"required"`
}
