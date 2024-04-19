package model

import (
	"time"
)

// UserTab [...]
type UserTab struct {
	ID         int64     `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null" json:"id"`
	Openid     string    `gorm:"index:idx_openid;column:openid;type:varchar(64);not null" json:"openid"`
	Money      int64     `gorm:"column:money;type:bigint;not null;default:0" json:"money"`
	GameTime   int64     `gorm:"column:game_time;type:bigint;not null;default:0" json:"gameTime"`
	State      int8      `gorm:"column:state;type:tinyint;not null;default:1;comment:'1-active 2-inactive'" json:"state"` // 1-active 2-inactive
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;default:null" json:"updateTime"`
	BizData    string    `gorm:"column:biz_data;type:mediumtext;default:null" json:"bizData"`
}

// UserTabColumns get sql column name.获取数据库列名
var UserTabColumns = struct {
	ID         string
	Openid     string
	Money      string
	GameTime   string
	State      string
	CreateTime string
	UpdateTime string
	BizData    string
}{
	ID:         "id",
	Openid:     "openid",
	Money:      "money",
	GameTime:   "game_time",
	State:      "state",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	BizData:    "biz_data",
}
