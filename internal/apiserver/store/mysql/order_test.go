package mysql

import (
	"context"
	"log"
	"os"
	"testing"

	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
	"github.com/teamen/kays/internal/pkg/options"
)

func TestOrderStoreInsCreate(t *testing.T) {

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
	defer mysqlStore.Close()
	var ds *datastore

	if v, ok := mysqlStore.(*datastore); ok {
		ds = v
	}

	log.Printf("%#v\n", ds)

	orderStore := newOrders(ds)
	customerStore := newCustomer(ds)
	ctx := context.Background()

	customer, err := customerStore.Get(ctx, 1)
	if err != nil {
		t.Fatalf("fail to get customer: %v", err)
	}

	var order v1.Order
	order.CustomerID = customer.ID
	order.Type = v1.ORDER_TYPE_PROCESS

	if err := orderStore.Create(ctx, &order); err != nil {
		t.Fatalf("failed to create order:%v", err)
	}

	t.Logf("%+v", order)

}
