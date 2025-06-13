package globals

import (
	"SchoolChat/database"
	"SchoolChat/utils"
)

var B *utils.BloomFilter

// InitBloomFilter 初始化布隆过滤器
func InitBloomFilter() {
	B = utils.InitBF(2000)
	users := make([]database.UserInfo, 0)
	database.Db.Find(&users)
	// 将所有用户名放入到布隆过滤器中
	for _, v := range users {
		B.SetItem(v.UID)
	}
}
