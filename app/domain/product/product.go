package product

import (
	"unicode/utf8"

	errDomain "github.com/radish-miyazaki/code-kakitai/domain/error"
	"github.com/radish-miyazaki/go-pkg/ulid"
)

const (
	nameLengthMin = 1
	nameLengthMax = 100

	descriptionLengthMin = 1
	descriptionLengthMax = 1000
)

type Product struct {
	id          string
	ownerID     string
	name        string
	description string
	price       int64
	stock       int
}

func newProduct(
	id string,
	ownerID string,
	name string,
	description string,
	price int64,
	stock int,
) (*Product, error) {
	if !ulid.IsValid(ownerID) {
		return nil, errDomain.NewError("owner id is invalid")
	}

	if utf8.RuneCountInString(name) < nameLengthMin || utf8.RuneCountInString(name) > nameLengthMax {
		return nil, errDomain.NewError("name is invalid")
	}

	if utf8.RuneCountInString(description) < descriptionLengthMin || utf8.RuneCountInString(description) > descriptionLengthMax {
		return nil, errDomain.NewError("description is invalid")
	}

	if price <= 0 {
		return nil, errDomain.NewError("price is invalid")
	}

	if stock < 0 {
		return nil, errDomain.NewError("stock is invalid")
	}

	return &Product{
		id:          id,
		ownerID:     ownerID,
		name:        name,
		description: description,
		price:       price,
		stock:       stock,
	}, nil
}

func Reconstruct(
	id string,
	ownerID string,
	name string,
	description string,
	price int64,
	stock int,
) (*Product, error) {
	return newProduct(id, ownerID, name, description, price, stock)
}

func NewProduct(
	ownerID string,
	name string,
	description string,
	price int64,
	stock int,
) (*Product, error) {
	return newProduct(ulid.NewULID(), ownerID, name, description, price, stock)
}

func (p *Product) ID() string {
	return p.id
}

func (p *Product) Price() int64 {
	return p.price
}

func (p *Product) Consume(quantity int) error {
	if quantity < 0 {
		return errDomain.NewError("product quantity is invalid")
	}

	if p.stock-quantity < 0 {
		return errDomain.NewError("stock is not enough")
	}

	p.stock -= quantity
	return nil
}
