package v1

import (
	"context"
	"log"

	"github.com/teamen/kays/internal/apiserver/store"
	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
	"github.com/teamen/kays/internal/pkg/util/idmaker"
)

type OrderSrv interface {
	Create(ctx context.Context, order *v1.Order) error
}

var _ OrderSrv = (*orderService)(nil)

type orderService struct {
	store store.Factory
}

func newOrders(srv *service) *orderService {
	return &orderService{store: srv.store}
}

func (o *orderService) Create(ctx context.Context, order *v1.Order) error {
	log.Printf("%+v", order)
	order.Serial = idmaker.GenId()

	if err := o.store.Orders().Create(ctx, order); err != nil {
		return err
	}
	return nil
}
