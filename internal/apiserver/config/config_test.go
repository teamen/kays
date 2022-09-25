package config

import (
	"testing"

	"github.com/teamen/kays/internal/apiserver/options"
	genericoptions "github.com/teamen/kays/internal/pkg/options"
)

func TestCreateConfigFromOptions(t *testing.T) {
	opts := &options.Options{
		MySQLOptions: genericoptions.NewMySQLOptions(),
	}
	config, _ := CreateConfigFromOptions(opts)
	t.Logf("%+v", config)
}
