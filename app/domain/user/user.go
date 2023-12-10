package user

import (
	"net/mail"
	"unicode/utf8"

	errDomain "github.com/radish-miyazaki/code-kakitai/domain/error"
	"github.com/radish-miyazaki/go-pkg/strings"
	"github.com/radish-miyazaki/go-pkg/ulid"
)

const (
	nameLengthMin = 1
	nameLengthMax = 255

	PhoneNumberDigitTen    = 10
	PhoneNumberDigitEleven = 11
)

var (
	phoneNumberDigitMap = map[int]struct{}{
		PhoneNumberDigitTen:    {},
		PhoneNumberDigitEleven: {},
	}
)

type Address struct {
	prefecture string
	city       string
	extra      string
}

func NewAddress(prefecture, city, extra string) (*Address, error) {
	if prefecture == "" || city == "" || extra == "" {
		return nil, errDomain.NewError("address value is invalid")
	}

	return &Address{
		prefecture: prefecture,
		city:       city,
		extra:      extra,
	}, nil
}

type User struct {
	id          string
	email       string
	phoneNumber string
	lastName    string
	firstName   string
	address     Address
}

func newUser(id, email, phoneNumber, lastName, firstName string, address Address) (*User, error) {
	if utf8.RuneCountInString(lastName) < nameLengthMin || utf8.RuneCountInString(lastName) > nameLengthMax {
		return nil, errDomain.NewError("lastName length is invalid")
	}

	if utf8.RuneCountInString(firstName) < nameLengthMin || utf8.RuneCountInString(firstName) > nameLengthMax {
		return nil, errDomain.NewError("firstName length is invalid")
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return nil, errDomain.NewError("email is invalid")
	}

	phoneNumber = strings.RemoveHyphen(phoneNumber)
	if _, ok := phoneNumberDigitMap[utf8.RuneCountInString(phoneNumber)]; !ok {
		return nil, errDomain.NewError("phoneNumber is invalid")
	}

	return &User{
		id:          id,
		email:       email,
		phoneNumber: phoneNumber,
		lastName:    lastName,
		firstName:   firstName,
		address:     address,
	}, nil
}

func Reconstruct(id, email, phoneNumber, lastName, firstName string, address Address) (*User, error) {
	return newUser(id, email, phoneNumber, lastName, firstName, address)
}

func NewUser(email, phoneNumber, lastName, firstName string, address Address) (*User, error) {
	return newUser(ulid.NewULID(), email, phoneNumber, lastName, firstName, address)
}

func (u *User) ID() string {
	return u.id
}

func (u *User) Email() string {
	return u.email
}
