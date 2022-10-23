package mysql

import (
	"context"

	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
	"gorm.io/gorm"
)

type orders struct {
	db *gorm.DB
}

func newOrders(ds *datastore) *orders {
	return &orders{db: ds.db}
}

func (o *orders) Create(ctx context.Context, order *v1.Order) error {
	return o.db.Save(order).Error
}

func (o *orders) Update(ctx context.Context, order *v1.Order) error {
	return o.db.Save(order).Error
}

func (o *orders) Get(ctx context.Context, ID uint32) (*v1.Order, error) {
	var order v1.Order
	if err := o.db.Model(&v1.Order{}).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
