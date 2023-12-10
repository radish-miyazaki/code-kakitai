package query_service

import (
	"context"

	"github.com/radish-miyazaki/code-kakitai/application/product"
	"github.com/radish-miyazaki/code-kakitai/infrastructure/mysql/db"
)

type productQueryService struct{}

func NewProductQueryService() *product.ProductQueryService {
	return &productQueryService{}
}

func (s *productQueryService) FetchProductList(ctx context.Context) ([]*product.FetchProductListDTO, error) {
	q := db.GetReadQuery(ctx)
	productWithOwners, err := q.ProductFetchWithOwner(ctx)
	if err != nil {
		return nil, err
	}

	var productFetchServiceDTOs []*product.FetchProductListDTO
	for _, productWithOwner := range productWithOwners {
		productFetchServiceDTOs = append(productFetchServiceDTOs, &product.FetchProductListDTO{
			ID:        productWithOwner.ID,
			Name:      productWithOwner.Name,
			Price:     productWithOwner.Price,
			Stock:     int(productWithOwner.Stock),
			OwnerID:   productWithOwner.OwnerID,
			OwnerName: productWithOwner.OwnerName.String,
		})
	}
	return productFetchServiceDTOs, nil
}
