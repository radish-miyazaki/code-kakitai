package owner

import (
	"net/mail"
	"unicode/utf8"

	errDomain "github.com/radish-miyazaki/code-kakitai/domain/error"
	"github.com/radish-miyazaki/go-pkg/ulid"
)

const (
	nameLengthMin = 1
	nameLengthMax = 255
)

type Owner struct {
	id    string
	name  string
	email string
}

func newOwner(id, name, email string) (*Owner, error) {
	if utf8.RuneCountInString(name) < nameLengthMin || nameLengthMax < utf8.RuneCountInString(name) {
		return nil, errDomain.NewError("name is invalid")
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return nil, errDomain.NewError("email is invalid")
	}

	return &Owner{
		id:    id,
		name:  name,
		email: email,
	}, nil
}

func NewOwner(name, email string) (*Owner, error) {
	return newOwner(ulid.NewULID(), name, email)
}

func Reconstruct(id, name, email string) (*Owner, error) {
	return newOwner(id, name, email)
}

func (o *Owner) ID() string {
	return o.id
}

func (o *Owner) Name() string {
	return o.name
}

func (o *Owner) Email() string {
	return o.email
}
