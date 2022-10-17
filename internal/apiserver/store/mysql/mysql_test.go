package mysql

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teamen/kays/internal/pkg/options"
)

func TestGetMySQLFactoryOr(t *testing.T) {

	opts := options.NewMySQLOptions()

	opts.Host = os.Getenv("MYSQL_HOST")
	opts.Username = os.Getenv("MYSQL_USER")
	opts.Password = os.Getenv("MYSQL_PWD")
	opts.Database = os.Getenv("MYSQL_DB")

	storeIns, err := GetMySQLFactoryOr(opts)
	defer storeIns.Close()

	assert.Empty(t, err)
	assert.NotEmpty(t, storeIns)

	storeIns2, _ := GetMySQLFactoryOr(nil)
	assert.Equal(t, storeIns, storeIns2)
}
