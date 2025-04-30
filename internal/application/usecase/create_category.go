package usecase

import (
	"context"

	"github.com/Vagmacker/luzora-api/internal/application/repository"
	"github.com/Vagmacker/luzora-api/internal/domain/category"
)

type (
	UseCase struct {
		categoryRepository repository.CategoryRepository
	}

	Input struct {
		Name string
	}
)

func New(categoryRepository repository.CategoryRepository) *UseCase {
	return &UseCase{
		categoryRepository: categoryRepository,
	}
}

func (usecase UseCase) Execute(ctx context.Context, input Input) error {
	category, err := category.New(input.Name)
	if err != nil {
		return err
	}

	err = usecase.categoryRepository.Save(ctx, category)
	if err != nil {
		return err
	}

	return nil
}
