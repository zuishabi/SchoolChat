package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var dsn = "root:861214959@tcp(127.0.0.1:3306)/game?charset=utf8mb4&parseTime=True&loc=Local"
var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&UserInfo{})
	Db.AutoMigrate(&ClassInfo{})
	Db.AutoMigrate(&ClassUser{})
	Db.AutoMigrate(&ChatBoard{})
}

type UserInfo struct {
	UID          string `gorm:"primarykey"`
	UserName     string `gorm:"Index:idx_name_psw"`
	Password     string `gorm:"Index:idx_name_psw"`
	GraduateTime string //毕业的年份
}

type ClassInfo struct {
	CID          uint32 `gorm:"primarykey"`
	CName        string
	CDescription string
	GraduateTime string `gorm:"Index"`
	OP           string `gorm:"Index"` //管理员id
}

// ClassUser 班级和用户的对应表
type ClassUser struct {
	CID      uint32 `gorm:"Index"`
	UID      string `gorm:"Index:idx_u"`
	UserName string `gorm:"Index:idx_u"`
}

// ChatBoard 用户的留言板
type ChatBoard struct {
	ID        uint32 `gorm:"primarykey"`
	UID       string `gorm:"Index"`
	Sender    string
	Content   string
	CreatedAt time.Time
}
