package order

import (
	"context"
	"time"

	cartDomain "github.com/radish-miyazaki/code-kakitai/domain/cart"
	errDomain "github.com/radish-miyazaki/code-kakitai/domain/error"
	productDomain "github.com/radish-miyazaki/code-kakitai/domain/product"
)

type orderDomainService struct {
	orderRepo   OrderRepository
	productRepo productDomain.ProductRepository
}

func NewOrderDomainService(orderRepo OrderRepository, productRepo productDomain.ProductRepository) OrderDomainService {
	return &orderDomainService{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

func (ds *orderDomainService) OrderProducts(ctx context.Context, cart *cartDomain.Cart, now time.Time) (string, error) {
	pa, err := ds.productRepo.FindByIDs(ctx, cart.ProductIDs())
	if err != nil {
		return "", err
	}

	productMap := make(map[string]*productDomain.Product)
	for _, p := range pa {
		productMap[p.ID()] = p
	}

	ops := make(OrderProducts, 0, len(cart.Products()))
	for _, cp := range cart.Products() {
		p, ok := productMap[cp.ProductID()]
		if !ok {
			return "", errDomain.NewError("product not found")
		}

		op, err := NewOrderProduct(p, cp.Quantity())
		if err != nil {
			return "", err
		}
		ops = append(ops, *op)

		if err := p.Consume(cp.Quantity()); err != nil {
			return "", err
		}
		if err := ds.productRepo.Save(ctx, p); err != nil {
			return "", err
		}
	}

	o, err := NewOrder(cart.UserID(), ops.TotalAmount(), ops, now)
	if err != nil {
		return "", err
	}
	if err = ds.orderRepo.Save(ctx, o); err != nil {
		return "", err
	}

	return "", nil
}
