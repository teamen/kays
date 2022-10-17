package mysql

import (
	"context"

	"gorm.io/gorm"

	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
	"github.com/teamen/kays/internal/pkg/util/gormutil"
	metav1 "github.com/teamen/kays/pkg/meta/v1"
)

const (
	FRAME_TYPE_FULL string = "full"
	FRAME_TYPE_HALF string = "half"
	FRAME_TYPE_LESS string = "less"
)

const (
	PRODUCT_STATUS_DISABLED int = iota + 0
	PRODUCT_STATUS_ACTIVE
)

type products struct {
	db *gorm.DB
}

func newProduct(ds *datastore) *products {
	return &products{ds.db}
}

func (p *products) Create(ctx context.Context, product *v1.Product) error {
	return p.db.Create(product).Error
}

func (p *products) Get(ctx context.Context, ID uint32) (*v1.Product, error) {

	var product v1.Product
	if err := p.db.Model(&v1.Product{}).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *products) List(ctx context.Context, opts metav1.ListOptions) (*v1.ProductList, error) {
	ret := &v1.ProductList{}

	lo := gormutil.ParseLimitAndOffset(opts.Page, opts.PageSize)
	res := p.db.
		Offset(lo.Offset).
		Limit(lo.Limit).
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)
	return ret, res.Error
}

func (p *products) Update(ctx context.Context, product *v1.Product) error {
	return p.db.Save(product).Error
}

func (p *products) Delete(ctx context.Context, ID uint32) error {
	return p.db.Delete(&v1.Product{}, "ID=?", ID).Error
}
