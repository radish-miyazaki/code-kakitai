package product

import (
	"context"

	productDomain "github.com/radish-miyazaki/code-kakitai/domain/product"
)

type SaveProductUseCase struct {
	productRepo productDomain.ProductRepository
}

func NewSaveProductUseCase(productRepo productDomain.ProductRepository) *SaveProductUseCase {
	return &SaveProductUseCase{
		productRepo: productRepo,
	}
}

type SaveProductUseCaseInputDTO struct {
	OwnerID     string
	Name        string
	Description string
	Price       int64
	Stock       int
}

type SaveProductUseCaseOutputDTO struct {
	ID          string
	OwnerID     string
	Name        string
	Description string
	Price       int64
	Stock       int
}

func (uc *SaveProductUseCase) Run(
	ctx context.Context,
	input SaveProductUseCaseInputDTO,
) (*SaveProductUseCaseOutputDTO, error) {
	p, err := productDomain.NewProduct(input.OwnerID, input.Name, input.Description, input.Price, input.Stock)
	if err != nil {
		return nil, err
	}

	err = uc.productRepo.Save(ctx, p)
	if err != nil {
		return nil, err
	}

	return &SaveProductUseCaseOutputDTO{
		ID:          p.ID(),
		OwnerID:     p.OwnerID(),
		Name:        p.Name(),
		Description: p.Description(),
		Price:       p.Price(),
		Stock:       p.Stock(),
	}, nil
}
