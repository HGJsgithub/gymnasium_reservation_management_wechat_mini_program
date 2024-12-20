package annCtrl

import (
	"Gymnasium_reservation_WeChat_mini_program/common/database"
	"Gymnasium_reservation_WeChat_mini_program/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

func DeleteAnn(c *gin.Context) {
	db := database.ConnectToGormDB()
	id := c.PostForm("id")
	annID, _ := strconv.Atoi(id)
	ann := model.Announcement{
		ID: annID,
	}
	db.Delete(&ann)
	c.JSON(http.StatusOK, gin.H{
		"msg": "成功删除！",
	})
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)

}
