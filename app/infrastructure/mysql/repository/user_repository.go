package repository

import (
	"context"
	"database/sql"
	"errors"

	errDomain "github.com/radish-miyazaki/code-kakitai/domain/error"
	userDomain "github.com/radish-miyazaki/code-kakitai/domain/user"
	"github.com/radish-miyazaki/code-kakitai/infrastructure/mysql/db"
)

type userRepository struct{}

func NewUserRepository() userDomain.UserRepository {
	return &userRepository{}
}

func (r *userRepository) Save(ctx context.Context, user *userDomain.User) error {
	return nil
}

func (r *userRepository) FindByID(ctx context.Context, id string) (*userDomain.User, error) {
	q := db.GetQuery(ctx)
	u, err := q.UserFindByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errDomain.ErrNotFound
		}

		return nil, err
	}

	uad, err := userDomain.NewAddress(u.Prefecture, u.City, u.AddressExtra)
	if err != nil {
		return nil, err
	}

	ud, err := userDomain.Reconstruct(
		u.ID,
		u.Email,
		u.PhoneNumber,
		u.LastName,
		u.FirstName,
		*uad,
	)
	if err != nil {
		return nil, err
	}

	return ud, nil
}

func (r *userRepository) FindAll(ctx context.Context) ([]*userDomain.User, error) {
	return nil, nil
}
