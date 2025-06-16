package globals

import (
	"SchoolChat/database"
	"gorm.io/gorm"
	"sync"
	"time"
)

type UpdateUserInfoReq struct {
	Name         string `json:"name"`
	GraduateTime string `json:"graduate_time"`
	Description  string `json:"description"`
}

var updateMap map[string]UpdateUserInfoReq
var lock sync.Mutex

func InitUpdater() {
	updateMap = make(map[string]UpdateUserInfoReq)
	go func() {
		for {
			select {
			case <-time.After(time.Second * 5):
				lock.Lock()
				for i, v := range updateMap {
					_ = database.Db.Transaction(func(tx *gorm.DB) error {
						tx.Model(database.ClassUser{}).Where("uid = ?", i).Updates(database.ClassUser{UserName: v.Name})
						tx.Model(database.EnterRequest{}).Where("uid = ?", i).Updates(database.EnterRequest{UserName: v.Name})
						tx.Model(database.ChatBoard{}).Where("sender = ?", i).Updates(database.ChatBoard{SenderName: v.Name})
						return nil
					})
				}
				clear(updateMap)
				lock.Unlock()
			}
		}
	}()
}

func AddUpdate(uid string, req UpdateUserInfoReq) {
	lock.Lock()
	defer lock.Unlock()
	updateMap[uid] = req
}
