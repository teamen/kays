package v1

import (
	"context"

	"github.com/teamen/kays/internal/apiserver/store"
	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
)

type UserSrv interface {
	Create(ctx context.Context, user *v1.User) error
	FindByUsername(ctx context.Context, usernname string) (*v1.User, error)
}

type userService struct {
	store store.Factory
}

// make sure that userService has been implement the UserSrv interface
var _ UserSrv = (*userService)(nil)

func newUsers(srv *service) *userService {
	return &userService{
		store: srv.store,
	}
}

func (u *userService) Create(ctx context.Context, user *v1.User) error {
	if err := u.store.Users().Create(ctx, user); err != nil {
		return err
	}
	return nil
}

func (u *userService) FindByUsername(ctx context.Context, usernname string) (*v1.User, error) {

	user, err := u.store.Users().FindByUsername(ctx, usernname)
	if err != nil {
		return nil, err
	}
	return user, nil
}
