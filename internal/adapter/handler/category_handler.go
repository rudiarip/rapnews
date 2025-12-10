package handler

import (
	"context"
	"rapnews/internal/core/domain/entity"
)

type CategoryHandler interface {
	GetCategories(ctx context.Context) ([]entity.CategoryEntity, error)
	GetCategoryByID(ctx context.Context, id int64) (*entity.CategoryEntity, error)
	CreateCategory(ctx context.Context, req entity.CategoryEntity) error
	EditCategoryByID(ctx context.Context, req entity.CategoryEntity) error
	DeleteCategoryByID(ctx context.Context, id int64) error
}
