package options

import (
	"encoding/json"
	"testing"

	"github.com/mitchellh/mapstructure"
)

func TestMapstructureRedisOptions(t *testing.T) {
	m := map[string]interface{}{
		"host":                     "redis",
		"port":                     6379,
		"database":                 10,
		"addrs":                    []string{},
		"username":                 "foo",
		"password":                 "bar",
		"master-name":              "MASTER",
		"optimisation-max-idle":    2000,
		"optimisation-max-active":  4000,
		"timeout":                  30,
		"use-ssl":                  false,
		"enable-cluster":           false,
		"ssl-insecure-skip-verify": false,
	}

	o := &RedisOptions{}
	mapstructure.Decode(m, o)
	t.Logf("%+v", o)

	bytes, _ := json.Marshal(o)
	t.Log(string(bytes))
	var opt RedisOptions
	json.Unmarshal(bytes, &opt)
	t.Logf("%+v", opt)
	t.Log(opt.String())
}
