package v1

import (
	"context"

	"github.com/teamen/kays/internal/apiserver/store"
	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
)

type CustomerSrv interface {
	Get(ctx context.Context, ID uint32) (*v1.Customer, error)
}

var _ CustomerSrv = (*customerService)(nil)

type customerService struct {
	store store.Factory
}

func newCustomers(srv *service) *customerService {
	return &customerService{store: srv.store}
}

func (customers *customerService) Get(ctx context.Context, ID uint32) (*v1.Customer, error) {
	c, err := customers.store.Customers().Get(ctx, ID)
	if err != nil {
		return nil, err
	}
	return c, nil
}
