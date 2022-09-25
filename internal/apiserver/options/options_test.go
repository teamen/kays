package options

import (
	"encoding/json"
	"testing"

	"github.com/mitchellh/mapstructure"
)

func TestNewOptions(t *testing.T) {
	o := NewOptions()
	// if bytes, err := json.Marshal(*mysqlOptions); err == nil {
	// 	// {"host":"127.0.0.1:3306","username":"","password":"","database":"","max-idle-connections":100,"max-open-connections":100,"max-connection-life-time":10000000000,"log-level":1}
	// 	fmt.Println(string(bytes))
	// }

	bytes, _ := json.Marshal(o)

	t.Logf("%+v", string(bytes))

}

func TestUnmarshal(t *testing.T) {
	str := `
	{
		"mysql": {
			"host": "127.0.0.1:3306",
			"database": "",
			"max-idle-connections": 100,
			"max-open-connections": 100,
			"max-connection-life-time": 10000000000,
			"log-level": 1
		},
		"redis": {
			"host": "127.0.0.1",
			"port": 6379,
			"addrs": [],
			"username": "",
			"password": "",
			"database": 0,
			"master-name": "",
			"optimisation-max-idle": 2000,
			"optimisation-max-active": 4000,
			"timeout": 0,
			"enable-cluster": false,
			"use-ssl": false,
			"ssl-insecure-skip-verify": false
		}
	}
	`

	o := &Options{}
	if err := json.Unmarshal([]byte(str), o); err != nil {
		t.Fatal(err.Error())
	}

	t.Logf("%s", o)

	t.Logf("%+v", o)

	var m map[string]interface{}
	if err := json.Unmarshal([]byte(str), &m); err != nil {
		t.Fatal(err.Error())
	}
	t.Log(m)
	var opt Options
	mapstructure.Decode(m, &opt)
	t.Log(opt.String())
}
