package model

type CategorySearch struct {
	Page     int
	PageSize int
}

type Category struct {
	ID          int
	Name        string
	Description string
}
