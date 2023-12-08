//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE
package order

import (
	"context"
	"time"

	cartDomain "github.com/radish-miyazaki/code-kakitai/domain/cart"
	errDomain "github.com/radish-miyazaki/code-kakitai/domain/error"
	productDomain "github.com/radish-miyazaki/code-kakitai/domain/product"
	"github.com/radish-miyazaki/go-pkg/ulid"
)

type OrderProduct struct {
	productID string
	price     int64
	quantity  int
}

func NewOrderProduct(product *productDomain.Product, quantity int) (*OrderProduct, error) {
	return &OrderProduct{
		productID: product.ID(),
		price:     product.Price(),
		quantity:  quantity,
	}, nil
}

func (op *OrderProduct) ProductID() string {
	return op.productID
}

func (op *OrderProduct) Price() int64 {
	return op.price
}

func (op *OrderProduct) Quantity() int {
	return op.quantity
}

type OrderProducts []OrderProduct

func (ops OrderProducts) ProductIDs() []string {
	productIDs := make([]string, 0, len(ops))
	for _, op := range ops {
		productIDs = append(productIDs, op.productID)
	}

	return productIDs
}

func (ops OrderProducts) TotalAmount() int64 {
	var totalAmount int64
	for _, op := range ops {
		totalAmount += op.price * int64(op.quantity)
	}

	return totalAmount
}

type Order struct {
	id          string
	userID      string
	totalAmount int64
	products    OrderProducts
	orderedAt   time.Time
}

func newOrder(id string, userID string, totalAmount int64, products OrderProducts, orderedAt time.Time) (*Order, error) {
	if !ulid.IsValid(id) {
		return nil, errDomain.NewError("id is invalid")
	}

	if totalAmount < 0 {
		return nil, errDomain.NewError("totalAmount is invalid")
	}

	if len(products) <= 0 {
		return nil, errDomain.NewError("products is nothing")
	}

	return &Order{
		id:          id,
		userID:      userID,
		totalAmount: totalAmount,
		products:    products,
		orderedAt:   orderedAt,
	}, nil
}

func NewOrder(userID string, totalAmount int64, products OrderProducts, orderedAt time.Time) (*Order, error) {
	return newOrder(
		ulid.NewULID(),
		userID,
		totalAmount,
		products,
		orderedAt,
	)
}

func Reconstruct(id string, userID string, totalAmount int64, products OrderProducts, orderedAt time.Time) (*Order, error) {
	return newOrder(
		id,
		userID,
		totalAmount,
		products,
		orderedAt,
	)
}

func (o *Order) ID() string {
	return o.id
}

func (o *Order) UserID() string {
	return o.userID
}

func (o *Order) TotalAmount() int64 {
	return o.totalAmount
}

func (o *Order) Products() OrderProducts {
	return o.products
}

func (o *Order) OrderedAt() time.Time {
	return o.orderedAt
}

type OrderDomainService interface {
	OrderProducts(ctx context.Context, cart *cartDomain.Cart, now time.Time) (string, error)
}
