package options

import (
	"encoding/json"

	genericoptions "github.com/teamen/kays/internal/pkg/options"
)

type Options struct {
	MySQLOptions *genericoptions.MySQLOptions `json:"mysql" mapstructure:"mysql"`
	RedisOptions *genericoptions.RedisOptions `json:"redis" mapstructure:"redis"`
}

func NewOptions() *Options {
	o := Options{
		MySQLOptions: genericoptions.NewMySQLOptions(),
		RedisOptions: genericoptions.NewRedisOptions(),
	}
	return &o
}

func (o *Options) String() string {
	bytes, _ := json.Marshal(o)
	return string(bytes)
}
