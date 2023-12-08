package cart

import (
	"time"

	errDomain "github.com/radish-miyazaki/code-kakitai/domain/error"
	"github.com/radish-miyazaki/go-pkg/ulid"
)

var (
	CartTimeout = time.Minute * 30
)

type CartProduct struct {
	productID string
	quantity  int
}

func NewCardProduct(productID string, quantity int) (*CartProduct, error) {
	if !ulid.IsValid(productID) {
		return nil, errDomain.NewError("productID is invalid")
	}

	if quantity <= 0 {
		return nil, errDomain.NewError("quantity is invalid")
	}

	return &CartProduct{
		productID: productID,
		quantity:  quantity,
	}, nil
}

func (cp *CartProduct) ProductID() string {
	return cp.productID
}

func (cp *CartProduct) Quantity() int {
	return cp.quantity
}

type CartProducts []CartProduct

type Cart struct {
	userID   string
	products CartProducts
}

func NewCart(userID string) (*Cart, error) {
	if !ulid.IsValid(userID) {
		return nil, errDomain.NewError("userID is invalid")
	}
	return &Cart{
		userID:   userID,
		products: CartProducts{},
	}, nil
}

func (c *Cart) UserID() string {
	return c.userID
}

func (c *Cart) Products() CartProducts {
	return c.products
}

func (c *Cart) ProductIDs() []string {
	ids := make([]string, 0, len(c.products))
	for _, p := range c.products {
		ids = append(ids, p.productID)
	}

	return ids
}

func (c *Cart) QuantityByProductID(productID string) (int, error) {
	for _, product := range c.products {
		if product.productID == productID {
			return product.quantity, nil
		}
	}

	return 0, errDomain.NewError("product not found")
}

func (c *Cart) AddProduct(productID string, quantity int) error {
	cp, err := NewCardProduct(productID, quantity)
	if err != nil {
		return err
	}

	for k, product := range c.products {
		if product.productID == productID {
			c.products[k] = *cp
			return nil
		}
	}

	c.products = append(c.products, *cp)
	return nil
}

func (c *Cart) RemoveProduct(productID string) error {
	var newProducts CartProducts
	for _, product := range c.products {
		if product.productID == productID {
			continue
		}

		newProducts = append(newProducts, product)
	}

	c.products = newProducts
	return nil
}
