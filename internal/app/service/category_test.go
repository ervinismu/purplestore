package service

import (
	"errors"
	"testing"

	"github.com/ervinismu/purplestore/internal/app/model"
	"github.com/ervinismu/purplestore/internal/app/schema"
	"github.com/ervinismu/purplestore/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCategoryService_Detail(t *testing.T) {
	type TestCase struct {
		Name        string
		Given       int
		Data        model.Category
		ExpectData  int
		ExpectError error
	}

	cases := []TestCase{
		{
			Name:  "when category data exist",
			Given: 1,
			Data: model.Category{
				ID:          1,
				Name:        "category 1",
				Description: "description category 1",
			},
			ExpectData:  1,
			ExpectError: nil,
		},
		{
			Name:  "when category data not exist",
			Given: 1,
			Data: model.Category{},
			ExpectData:  0,
			ExpectError: nil,
		},
		{
			Name:  "when error get category data",
			Given: 1,
			Data: model.Category{},
			ExpectData:  0,
			ExpectError: errors.New("error test"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			mockCategoryRepo := mocks.NewMockCategoryRepository(mockCtrl)
			mockCategoryRepo.
				EXPECT().
				GetByID(tc.Given).
				Return(tc.Data, tc.ExpectError)

			categoryService := NewCategorySerivce(mockCategoryRepo)

			req := schema.CategoryDetailRequest{ID: tc.Given}
			category, err := categoryService.Detail(req)

			assert.Equal(t, tc.ExpectData, category.ID)
			assert.Equal(t, tc.ExpectError, err)
		})
	}
}
