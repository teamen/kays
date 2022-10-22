package mysql

import (
	"log"
	"os"
	"testing"

	"github.com/teamen/kays/internal/apiserver/store"
	"github.com/teamen/kays/internal/pkg/options"
)

var mysqlStore store.Factory
var ds *datastore

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

	if v, ok := mysqlStore.(*datastore); ok {
		ds = v
	}
	t.Logf("%+v", ds)

	return func(t *testing.T) {
		log.Printf("close db connection...")
		mysqlStore.Close()
	}
}
