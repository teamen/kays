package v1

import (
	"context"
	"fmt"

	"github.com/teamen/kays/internal/apiserver/store"
	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
	metav1 "github.com/teamen/kays/pkg/meta/v1"
)

type ProductSrv interface {
	Create(ctx context.Context, product *v1.Product) error
	Get(ctx context.Context, ID uint32) (*v1.Product, error)
	Update(ctx context.Context, product *v1.Product) error
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ProductList, error)
	Delete(ctx context.Context, ID uint32) error
}

type productService struct {
	store store.Factory
}

func newProducts(srv *service) *productService {
	return &productService{store: srv.store}
}
func (products *productService) Create(ctx context.Context, product *v1.Product) error {
	if err := products.store.Products().Create(ctx, product); err != nil {
		return err
	}
	return nil
}

func (products *productService) Get(ctx context.Context, ID uint32) (*v1.Product, error) {
	product, err := products.store.Products().Get(ctx, ID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (products *productService) Update(ctx context.Context, product *v1.Product) error {
	err := products.store.Products().Update(ctx, product)
	return err
}

func (products *productService) List(ctx context.Context, opts metav1.ListOptions) (*v1.ProductList, error) {
	productList, err := products.store.Products().List(ctx, opts)
	if err != nil {
		return nil, err
	}
	fmt.Println(products)
	return productList, nil
}

func (products *productService) Delete(ctx context.Context, ID uint32) error {
	return products.store.Products().Delete(ctx, ID)
}
