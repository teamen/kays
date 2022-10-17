package mysql

import (
	"context"

	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
	"github.com/teamen/kays/internal/pkg/util/gormutil"
	metav1 "github.com/teamen/kays/pkg/meta/v1"
	"gorm.io/gorm"
)

const (
	CUSTOMER_STATUS_DISABLED int = iota + 0
	CUSTOMER_STATUS_ACTIVE
)

type customers struct {
	db *gorm.DB
}

func newCustomer(ds *datastore) *customers {
	return &customers{
		db: ds.db,
	}
}

func (c *customers) Create(ctx context.Context, customer *v1.Customer) error {
	return c.db.Create(customer).Error
}

func (c *customers) Get(ctx context.Context, ID uint32) (*v1.Customer, error) {
	var customer v1.Customer
	if err := c.db.Model(&v1.Customer{}).First(&customer, "ID=?", ID).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (c *customers) List(ctx context.Context, opts metav1.ListOptions) (*v1.CustomerList, error) {
	ret := &v1.CustomerList{}

	lo := gormutil.ParseLimitAndOffset(opts.Page, opts.PageSize)
	res := c.db.
		Offset(lo.Offset).
		Limit(lo.Limit).
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)
	return ret, res.Error
}

func (c *customers) Update(ctx context.Context, customer *v1.Customer) error {
	return c.db.Save(customer).Error
}

func (c *customers) Delete(ctx context.Context, ID uint32) error {
	err := c.db.Where("id = ?", ID).Delete(&v1.Customer{}).Error
	return err
}
