package options

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/mitchellh/mapstructure"
)

func TestNewMySQLOptions(t *testing.T) {
	mysqlOptions := NewMySQLOptions()
	// t.Logf("%+v", mysqlOptions)

	if bytes, err := json.Marshal(*mysqlOptions); err == nil {
		// {"host":"127.0.0.1:3306","username":"","password":"","database":"","max-idle-connections":100,"max-open-connections":100,"max-connection-life-time":10000000000,"log-level":1}
		fmt.Println(string(bytes))
	}
}

func TestMapstructureMySQLOptions(t *testing.T) {
	str := `{
		"host": "0.0.0.0:3306",
		"username": "root",
		"password": "root",
		"database": "",
		"max-idle-connections": 60,
		"max-open-connections": 60,
		"max-connection-life-time": 10000000000,
		"log-level": 4
	}`

	opt := NewMySQLOptions()
	if err := json.Unmarshal([]byte(str), &opt); err != nil {
		t.Fatalf(err.Error())
	}

	fmt.Println(opt.Host)
	bytes, _ := json.Marshal(&opt)
	fmt.Println(string(bytes))

	m := map[string]interface{}{
		"host":                     "127.0.0.1:3306",
		"username":                 "homestead",
		"password":                 "secret",
		"database":                 "kaysserver",
		"max-idle-connections":     100,
		"max-openn-connections":    100,
		"max-connection-life-time": 20000000000,
		"log-level":                3,
	}
	o := MySQLOptions{}
	mapstructure.Decode(m, &o)
	t.Logf("%+v", o)
	t.Log(o.String())

}
