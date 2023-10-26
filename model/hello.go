package model

import (
	"time"
)

// TStudent [...]
type TStudent struct {
	ID         int       `gorm:"autoIncrement:true;primaryKey;column:id;type:int;not null" json:"id"`
	Name       string    `gorm:"column:name;type:varchar(50);not null" json:"name"`
	Age        int       `gorm:"column:age;type:int;not null" json:"age"`
	Phone      string    `gorm:"column:phone;type:varchar(50);not null" json:"phone"`
	Cls        string    `gorm:"column:cls;type:varchar(50);not null" json:"cls"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;default:null" json:"updateTime"`
}

// TStudentColumns get sql column name.获取数据库列名
var TStudentColumns = struct {
	ID         string
	Name       string
	Age        string
	Phone      string
	Cls        string
	CreateTime string
	UpdateTime string
}{
	ID:         "id",
	Name:       "name",
	Age:        "age",
	Phone:      "phone",
	Cls:        "cls",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// User [...]
type User struct {
	ID    int64  `gorm:"primaryKey;column:id;type:bigint;not null;comment:'主键ID'" json:"id"`   // 主键ID
	Name  string `gorm:"column:name;type:varchar(30);default:null;comment:'姓名'" json:"name"`   // 姓名
	Age   int    `gorm:"column:age;type:int;default:null;comment:'年龄'" json:"age"`             // 年龄
	Email string `gorm:"column:email;type:varchar(50);default:null;comment:'邮箱'" json:"email"` // 邮箱
}

// UserColumns get sql column name.获取数据库列名
var UserColumns = struct {
	ID    string
	Name  string
	Age   string
	Email string
}{
	ID:    "id",
	Name:  "name",
	Age:   "age",
	Email: "email",
}
