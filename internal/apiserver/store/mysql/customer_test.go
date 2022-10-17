package mysql

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/teamen/kays/internal/apiserver/store"
	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
	"github.com/teamen/kays/internal/pkg/options"
	metav1 "github.com/teamen/kays/pkg/meta/v1"
)

var mysqlStore store.Factory

func setUp(t *testing.T) func(t *testing.T) {
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

	return func(t *testing.T) {
		log.Printf("close db connection...")
		mysqlStore.Close()
	}
}

func TestNewCustomer(t *testing.T) {
	tearDown := setUp(t)
	defer tearDown(t)

	var ds *datastore

	if v, ok := mysqlStore.(*datastore); ok {
		ds = v
	}

	log.Printf("%+v\n", ds)
	log.Printf("%#v\n", ds)
	c := newCustomer(ds)

	log.Printf("%#v", c)

	customer := &v1.Customer{
		Name:    "Foo",
		Status:  CUSTOMER_STATUS_ACTIVE,
		Balance: 9999,
	}
	ctx := context.Background()
	if err := c.Create(ctx, customer); err != nil {
		t.Fatalf("failed to create customer")
	}

	log.Printf("customer created:%+v", customer)

	ID := customer.ID

	c2, err := c.Get(ctx, customer.ID)
	if err != nil {
		t.Logf("failed to find the customer[#%d]", customer.ID)
	} else {
		t.Logf("customer[#%d] found: %+v", customer.ID, c2)
	}

	c2.Name = "foobar"
	if err := c.Update(ctx, c2); err != nil {
		t.Log("failed to update")
	}

	c2, _ = c.Get(ctx, c2.ID)
	t.Logf("name updated: %s", c2.Name)

	if err := c.Delete(ctx, customer.ID); err != nil {
		t.Fatalf("failed to delete customer")
	}

	_, err = c.Get(ctx, ID)
	t.Logf("%+v\n", err)
	// t.Logf("%s\n", err.Error())

	var Page, PageSize int64 = 1, 30
	lo := &metav1.ListOptions{
		Page:     &Page,
		PageSize: &PageSize,
	}
	if customerList, err := c.List(ctx, *lo); err != nil {
		t.Logf("failed to list customers")
	} else {
		log.Printf("%+v", customerList)

		for i, perCustomer := range customerList.Items {
			log.Printf("[%d]: %+v\n", i, perCustomer)
		}
	}

}
