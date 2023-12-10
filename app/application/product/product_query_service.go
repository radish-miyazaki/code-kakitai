package product

import "context"

type FetchProductListDTO struct {
	ID        string
	Name      string
	Price     int64
	Stock     int
	OwnerID   string
	OwnerName string
}

type ProductQueryService interface {
	FetchProductList(ctx context.Context) ([]*FetchProductListDTO, error)
}
