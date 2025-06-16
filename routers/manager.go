package routers

import (
	"SchoolChat/database"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func ManageClass(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("uid") == nil {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	uid := session.Get("uid").(string)
	t := c.Query("cid")
	cid, _ := strconv.Atoi(t)
	// 检查当前用户是否是当前班级的管理员
	if !checkOP(uint32(cid), uid) {
		c.Redirect(http.StatusSeeOther, "/main")
		return
	}
	c.HTML(http.StatusOK, "manage.html", nil)
}

// DeleteClass 解散班级
func DeleteClass(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("uid") == nil {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	uid := session.Get("uid").(string)
	t := c.Query("cid")
	cid, _ := strconv.Atoi(t)
	if !checkOP(uint32(cid), uid) {
		c.Redirect(http.StatusSeeOther, "/main")
		return
	}
	err := database.Db.Transaction(func(tx *gorm.DB) error {
		u := database.ClassUser{}
		if err := tx.Where("c_id = ?", cid).Delete(&u).Error; err != nil {
			return err
		}
		return tx.Where("c_id = ?", cid).Delete(&database.ClassInfo{}).Error
	})
	if err != nil {
		c.String(http.StatusForbidden, err.Error())
		return
	}
	c.String(http.StatusOK, "解散成功")
}

// RemoveUser 移除班级中的用户
func RemoveUser(c *gin.Context) {
	uid := c.Query("uid")
	cid, _ := strconv.Atoi(c.Query("cid"))
	if !checkOP(uint32(cid), uid) {
		c.String(http.StatusForbidden, "权限不足")
		return
	}
	database.Db.Where("uid = ?", uid).Delete(&database.ClassUser{})
	c.String(http.StatusOK, "删除成功")
}

// GetEnterRequest 获得所有的加入请求
func GetEnterRequest(c *gin.Context) {
	cid, _ := strconv.Atoi(c.Query("cid"))
	users := make([]database.EnterRequest, 0)
	database.Db.Where("c_id = ?", cid).Find(&users)
	res := make([]Member, len(users))
	for i, v := range users {
		res[i].ID = v.UID
		res[i].Name = v.UserName
	}
	c.JSON(http.StatusOK, res)
}

// ApproveEnter 同意加入班级
func ApproveEnter(c *gin.Context) {
	uid := c.Query("uid")
	cid, _ := strconv.Atoi(c.Query("cid"))
	// 开启事务，将用户从加入请求表中删除，并加入用户列表
	err := database.Db.Transaction(func(tx *gorm.DB) error {
		// 判断是否有
		if err := tx.Where("uid = ?", uid).First(&database.ClassUser{}).Error; err == nil {
			return errors.New("当前用户已加入其他班级")
		}
		u := &database.EnterRequest{}
		tx.Where("uid = ?", uid).First(&u)
		tx.Where("uid = ?", uid).Delete(&database.EnterRequest{})
		// 将用户添加到班级表中
		user := database.ClassUser{
			CID:      uint32(cid),
			UID:      uid,
			UserName: u.UserName,
		}
		tx.Create(&user)
		return nil
	})
	if err != nil {
		c.String(http.StatusInternalServerError, "同意失败: %v", err)
		return
	}
	c.String(http.StatusOK, "")
}

// RejectEnter 拒绝加入班级
func RejectEnter(c *gin.Context) {
	uid := c.Query("uid")
	cid, _ := strconv.Atoi(c.Query("cid"))
	database.Db.Where("uid = ? AND cid = ?", uid, cid).Delete(&database.EnterRequest{})
	c.String(http.StatusOK, "拒绝成功")
}

// 检查当前用户是否是班级的管理员
func checkOP(cid uint32, uid string) bool {
	info := database.ClassInfo{}
	if err := database.Db.Where("c_id = ?", cid).First(&info).Error; err != nil {
		return false
	}
	if info.OP != uid {
		return false
	}
	return true
}
