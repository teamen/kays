package v1

import (
	"database/sql"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestMarshalUser(t *testing.T) {
	ipAddress := "192.168.0.1"
	ip := net.ParseIP(ipAddress)

	email := "waynetsecn@gmail.com"
	user := &User{
		ID:            1,
		Username:      "Wayne Tse",
		Email:         &email,
		Password:      "",
		RememberToken: "",
		LastLoginIp:   uint32(binary.BigEndian.Uint32(ip.To4())),
		LastLoginAt:   time.Now(),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		DeletedAt:     sql.NullTime{Valid: false},
	}
	bytes, _ := json.Marshal(user)
	fmt.Println(string(bytes))
}
