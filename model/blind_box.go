package model

import (
	"time"
)

// OrderTab [...]
type OrderTab struct {
	OrderID         int64     `gorm:"primaryKey;column:order_id;type:bigint;not null" json:"orderId"`
	Seller          string    `gorm:"index:idx_seller;column:seller;type:varchar(256);not null" json:"seller"`
	NftAddr         string    `gorm:"index:idx_nft_addr_token_id;column:nft_addr;type:varchar(256);not null" json:"nftAddr"`
	TokenID         int64     `gorm:"index:idx_nft_addr_token_id;column:token_id;type:bigint;not null" json:"tokenId"`
	Price           float64   `gorm:"column:price;type:decimal(27,0);not null" json:"price"`
	Amount          int64     `gorm:"column:amount;type:bigint;not null" json:"amount"`
	State           int8      `gorm:"column:state;type:tinyint;not null;default:1;comment:'1-active 2-inactive'" json:"state"`                             // 1-active 2-inactive
	CreateTimestamp int64     `gorm:"column:create_timestamp;type:bigint;not null;comment:'row created block timestamp'" json:"createTimestamp"`           // row created block timestamp
	UpdateTimestamp int64     `gorm:"column:update_timestamp;type:bigint;not null;default:0;comment:'row updated block timestamp'" json:"updateTimestamp"` // row updated block timestamp
	CreateTime      time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'db row create time'" json:"createTime"`  // db row create time
	UpdateTime      time.Time `gorm:"column:update_time;type:datetime;default:null;comment:'db row update time'" json:"updateTime"`                        // db row update time
}

// OrderTabColumns get sql column name.获取数据库列名
var OrderTabColumns = struct {
	OrderID         string
	Seller          string
	NftAddr         string
	TokenID         string
	Price           string
	Amount          string
	State           string
	CreateTimestamp string
	UpdateTimestamp string
	CreateTime      string
	UpdateTime      string
}{
	OrderID:         "order_id",
	Seller:          "seller",
	NftAddr:         "nft_addr",
	TokenID:         "token_id",
	Price:           "price",
	Amount:          "amount",
	State:           "state",
	CreateTimestamp: "create_timestamp",
	UpdateTimestamp: "update_timestamp",
	CreateTime:      "create_time",
	UpdateTime:      "update_time",
}
