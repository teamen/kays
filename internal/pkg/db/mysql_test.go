package db

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	opts := &Options{
		Host:                  "mysql:3306",
		Username:              os.Getenv("MYSQL_USER"),
		Password:              os.Getenv("MYSQL_PWD"),
		Database:              "demo",
		MaxIdleConnections:    400,
		MaxOpenConnections:    200,
		MaxConnectionLifeTime: time.Duration(10) * time.Second,
		LogLevel:              1,
	}
	// fmt.Printf("%+v", opts)
	db, err := New(opts)
	if err != nil {
		t.Fatalf("fail to connect db:%s", err.Error())
	}

	fmt.Println(db)
}
