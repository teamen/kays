package mysql

import (
	"context"

	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
	"gorm.io/gorm"
)

type orderProducts struct {
	db *gorm.DB
}

func newOrderProducts(ds *datastore) *orderProducts {
	return &orderProducts{ds.db}
}

func (op *orderProducts) Create(ctx context.Context, orderProduct *v1.OrderProduct) error {
	return nil
}

func (op *orderProducts) Delete(ctx context.Context, orderProductID uint32) error {
	return nil
}
