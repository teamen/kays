package v1

import (
	"context"
	"fmt"

	"github.com/teamen/kays/internal/apiserver/store"
	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
)

type UserSrv interface {
	Create(ctx context.Context, user *v1.User) error
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

	fmt.Printf("%+v", user)
	if err := u.store.Users().Create(ctx, user); err != nil {
		return err
	}
	return nil
}
