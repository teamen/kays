package v1

import (
	"database/sql"
	"time"
)

const (
	ORDER_TYPE_PROCESS = "process"
	ORDER_TYPE_REPAIR  = "repair"
)

type Order struct {
	ID              uint32       `gorm:"primaryKey;autoIncrement;comment:ID"`
	Serial          string       `gorm:"size:32;not null;comment:序号"`
	CustomerID      uint32       `gorm:"not null;default:0;comment:客户ID"`
	Type            string       `gorm:"type:char(16) not null;default:'process';comment:类别[process|repair]"`
	CustomerSerial  string       `gorm:"size:32;not null;default:'';comment:客户订单号"`
	DiscountPercent uint16       `gorm:"default:10000;comment:镜片费用折扣"`
	Owner           string       `gorm:"size:32;not null;default:'';comment:订单属主"`
	OwnerPhone      string       `gorm:"size:32;not null;default:'';comment:订单属主手机"`
	Status          uint8        `gorm:"type:tinyint(4);default:1;not null;comment:订单状态"`
	PayStatus       uint8        `gorm:"type:tinyint(4);not null;default:0;comment:支付状态"`
	Memo            string       `gorm:"size:255;not null;default:'';comment:备注"`
	AdjustFee       int32        `gorm:"not null;default:0;comment:差价"`
	MaterialFee     uint32       `gorm:"not null;default:0;comment:材料费"`
	ProcessFee      uint32       `gorm:"not null;default:0;comment:加工费"`
	TotalFee        uint32       `gorm:"not null;default:0;comment:总费用"`
	SettledFee      uint32       `gorm:"not null;default:0;comment:结算费用，折扣后"`
	CreatedAt       time.Time    `gorm:"column:created_at;type:timestamp null;" json:"created_at"`
	UpdatedAt       time.Time    `gorm:"column:updated_at;type:timestamp null;" json:"updated_at"`
	DeletedAt       sql.NullTime `gorm:"column:deleted_at;type:timestamp null;" json:"-,omitempty"`
	SettlementID    uint32       `gorm:"not null;default:0;comment:余额结算记录ID"`
}
