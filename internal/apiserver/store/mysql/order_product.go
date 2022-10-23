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
	if err := op.db.Create(orderProduct).Error; err != nil {
		return err
	}
	return nil
}

func (op *orderProducts) Delete(ctx context.Context, orderProductID uint32) error {
	return op.db.Delete(&v1.OrderProduct{}, "id=?", orderProductID).Error
}

func (op *orderProducts) List(ctx context.Context, condition map[string]interface{}) ([]*v1.OrderProduct, error) {

	var ret []*v1.OrderProduct
	if err := op.db.Model(&v1.OrderProduct{}).Where(condition).Find(&ret).Error; err != nil {
		return nil, err
	}
	return ret, nil
}
