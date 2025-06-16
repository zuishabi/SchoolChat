package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var dsn = "root:861214959@tcp(127.0.0.1:3309)/game?charset=utf8mb4&parseTime=True&loc=Local"
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
	Db.AutoMigrate(&EnterRequest{})
}

type UserInfo struct {
	UID          string `gorm:"primarykey"`
	UserName     string `gorm:"Index:idx_name_graduate"`
	Password     string
	GraduateTime string `gorm:"Index:idx_name_graduate"`
	Description  string
}

type ClassInfo struct {
	CID          uint32 `gorm:"primarykey"`
	CName        string
	CDescription string
	GraduateTime string `gorm:"Index"`
	OP           string //管理员id
}

// ClassUser 班级和用户的对应表
type ClassUser struct {
	CID      uint32 `gorm:"Index"`
	UID      string `gorm:"primarykey"`
	UserName string
}

// ChatBoard 用户的留言板
type ChatBoard struct {
	ID         uint32 `gorm:"primarykey"`
	UID        string `gorm:"Index"`
	SenderName string
	Sender     string
	Content    string
	CreatedAt  time.Time
}

// EnterRequest 用户加入班级的请求
type EnterRequest struct {
	CID      uint32 `gorm:"Index:idx_u"`
	UID      string `gorm:"Index:idx_u"`
	UserName string `gorm:"Index:idx_u"`
}
