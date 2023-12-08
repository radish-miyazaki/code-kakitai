package repository

import (
	"context"

	productDomain "github.com/radish-miyazaki/code-kakitai/domain/product"
)

type productRepository struct{}

func NewProductRepository() productDomain.ProductRepository {
	return &productRepository{}
}

func (r *productRepository) Save(ctx context.Context, product *productDomain.Product) error {
	return nil
}

func (r *productRepository) FindByID(ctx context.Context, id string) (*productDomain.Product, error) {
	return nil, nil
}

func (r *productRepository) FindByIDs(ctx context.Context, ids []string) ([]*productDomain.Product, error) {
	return nil, nil
}
