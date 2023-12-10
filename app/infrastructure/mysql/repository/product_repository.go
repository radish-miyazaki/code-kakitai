package repository

import (
	"context"

	productDomain "github.com/radish-miyazaki/code-kakitai/domain/product"
	"github.com/radish-miyazaki/code-kakitai/infrastructure/mysql/db"
	"github.com/radish-miyazaki/code-kakitai/infrastructure/mysql/db/db_gen"
)

type productRepository struct{}

func NewProductRepository() productDomain.ProductRepository {
	return &productRepository{}
}

func (r *productRepository) Save(ctx context.Context, product *productDomain.Product) error {
	query := db.GetQuery(ctx)
	if err := query.UpsertProduct(ctx, db_gen.UpsertProductParams{
		ID:          product.ID(),
		OwnerID:     product.OwnerID(),
		Name:        product.Name(),
		Description: product.Description(),
		Price:       product.Price(),
		Stock:       int32(product.Stock()),
	}); err != nil {
		return err
	}

	return nil
}

func (r *productRepository) FindByID(ctx context.Context, id string) (*productDomain.Product, error) {
	return nil, nil
}

func (r *productRepository) FindByIDs(ctx context.Context, ids []string) ([]*productDomain.Product, error) {
	return nil, nil
}
