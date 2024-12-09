package model

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	ID         string `gorm:"type:char(19)" json:"id"`
	Nickname   string `gorm:"size:10" json:"nickname"`
	Phone      string `gorm:"type:char(11);primary_key" json:"phone"`
	Password   string `gorm:"size:16" json:"password"`
	CreateTime string `gorm:"type:varchar(12)" json:"createTime"`
}

func (u *User) Avatar(c *gin.Context) {

}
