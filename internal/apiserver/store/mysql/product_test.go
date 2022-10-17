package mysql

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
	"github.com/teamen/kays/internal/pkg/options"
	metav1 "github.com/teamen/kays/pkg/meta/v1"
)

func TestProductCURD(t *testing.T) {

	opts := options.NewMySQLOptions()

	opts.Host = os.Getenv("MYSQL_HOST")
	opts.Username = os.Getenv("MYSQL_USER")
	opts.Password = os.Getenv("MYSQL_PWD")
	opts.Database = os.Getenv("MYSQL_DB")

	var err error
	mysqlStore, err = GetMySQLFactoryOr(opts)

	if err != nil {
		t.Fatalf("failed to connect mysql")
	}

	t.Logf("connected:%+v", mysqlStore)

	defer mysqlStore.Close()

	var ds *datastore

	if v, ok := mysqlStore.(*datastore); ok {
		ds = v
	}

	log.Printf("%+v\n", ds)
	log.Printf("%#v\n", ds)

	productStore := newProduct(ds)
	ctx := context.Background()

	product := &v1.Product{
		Name:             "L8-黑色413",
		Fullname:         "L8-413-黑色--L系列-TAPOLE-镜架",
		Code:             "413#",
		TypeCategoryID:   2,
		SeriesCategoryID: 0,
		FrameType:        FRAME_TYPE_HALF,
		Status:           int8(PRODUCT_STATUS_ACTIVE),
		Stock:            20,
		MaterialFee:      100,
		ExtraFee:         200,
	}
	if err := productStore.Create(ctx, product); err != nil {
		t.Fatalf("failed to create product:%s", err.Error())
	} else {
		t.Logf("product[%d] created:%#v", product.ID, product)
	}

	product.CustomerID = 1
	if err = productStore.Update(ctx, product); err != nil {
		t.Fatalf("failed to update product:%s", err.Error())
	} else {
		productUpdated, _ := productStore.Get(ctx, product.ID)
		t.Logf("product updated: %#v\n", productUpdated)
	}

	var productList *v1.ProductList
	var page, pageSize int64 = 1, 100
	listOptions := metav1.ListOptions{
		Page:     &page,
		PageSize: &pageSize,
	}

	productList, err = productStore.List(ctx, listOptions)
	if err != nil {
		t.Logf("failed to list products:%s", err.Error())
	}

	t.Logf("%#v", productList)
	t.Logf("%+v", productList.Items)

	for i, singleProduct := range productList.Items {
		t.Logf("deleting %d product #%d", i+1, singleProduct.ID)
		productStore.Delete(ctx, singleProduct.ID)
	}

	productList, _ = productStore.List(ctx, listOptions)
	assert.Equal(t, productList.TotalCount, int64(0))

}
