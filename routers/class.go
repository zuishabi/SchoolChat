package routers

import (
	"SchoolChat/database"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// CreateClass 创建一个班级
func CreateClass(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("uid") == nil {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	uid := session.Get("uid").(string)
	userName := session.Get("user_name").(string)
	u := database.ClassUser{}
	if err := database.Db.Where("uid = ?", uid).First(&u).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		// 用户没有加入班级，可以创建班级
		className := c.Query("class_name")
		classDescription := c.Query("class_description")

		err := database.Db.Transaction(func(tx *gorm.DB) error {
			class := database.ClassInfo{
				CName:        className,
				CDescription: classDescription,
				GraduateTime: session.Get("graduate_time").(string),
				OP:           uid,
			}
			if err := tx.Create(&class).Error; err != nil {
				return err
			}
			cu := database.ClassUser{
				CID:      class.CID,
				UID:      uid,
				UserName: userName,
			}
			if err := tx.Create(&cu).Error; err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			c.String(http.StatusInternalServerError, "创建班级失败: %v", err)
			return
		}
	} else {
		c.String(http.StatusInternalServerError, "您已加入班级")
	}
}

type classInfo struct {
	CID          uint32 `json:"cid"`
	CName        string `json:"c_name"`
	CDescription string `json:"c_description"`
	GraduateTime string `json:"graduate_time"`
	OP           string `json:"op"` //管理员id
}

// GetClass 返回当前加入的班级，并传递相关信息
func GetClass(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("uid") == nil {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	uid := session.Get("uid").(string)
	u := database.ClassUser{}
	// 查找当前用户是否加入了班级
	if err := database.Db.Where("uid = ?", uid).First(&u).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		c.String(http.StatusForbidden, "您未加入班级")
		return
	} else {
		info := classInfo{}
		temp := database.ClassInfo{}
		database.Db.Where("c_id = ?", u.CID).First(&temp)
		info.CID = temp.CID
		info.GraduateTime = temp.GraduateTime
		info.CDescription = temp.CDescription
		info.CName = temp.CName
		info.OP = temp.OP
		c.JSON(http.StatusOK, &info)
		return
	}
}

type Member struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetClassMembers 获取成员列表
func GetClassMembers(c *gin.Context) {
	cid := c.Query("cid")
	m := make([]database.ClassUser, 0)
	database.Db.Where("c_id = ?", cid).Find(&m)
	members := make([]Member, len(m))
	for i, v := range m {
		members[i].ID = v.UID
		members[i].Name = v.UserName
	}
	c.JSON(http.StatusOK, members)
}
