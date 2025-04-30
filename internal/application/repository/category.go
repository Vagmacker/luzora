package repository

import (
	"context"

	"github.com/Vagmacker/luzora-api/internal/domain/category"
)

type CategoryRepository interface {
	Save(ctx context.Context, c *category.Category) error
	Get(ctx context.Context, categoryId string) (category.Category, error)
}
