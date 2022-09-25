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

// import (
// 	"fmt"
// 	"time"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// type Options struct {
// 	MaxConnectionLifeTime time.Duration
// 	LogLevel              int
// 	Logger                logger.Interface
// }

// // New create a new gorm db instance with the given options.
// func New(opts *Options) (*gorm.DB, error) {

// 	dsn := fmt.Sprintf(
// 		`%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s`,
// 		opts.Username,
// 		opts.Password,
// 		opts.Host,
// 		opts.Database,
// 		true,
// 		"Local")

// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
// 		Logger: opts.Logger,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		return nil, err
// 	}

// 	sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)
// 	sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)
// 	sqlDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)

// 	return nil, nil
// }
