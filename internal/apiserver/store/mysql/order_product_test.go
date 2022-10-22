package mysql

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/teamen/kays/internal/pkg/options"
)

func TestOrderProductStoreInsCreate(t *testing.T) {

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

	ctx := context.Background()
	orderID := 1
	order, err := orderStore.Get(ctx, uint32(orderID))
	if err != nil {
		t.Logf("failed to retrive order #%d", orderID)
	}

	t.Logf("%+v", order)

	// orderProduct := &v1.OrderProduct{
	// 	CustomerID: order.CustomerID,
	// 	OrderID: order.ID,
	// }
}
