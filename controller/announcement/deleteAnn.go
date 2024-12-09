package announcement

import (
	"Gymnasium_reservation_WeChat_mini_program/common/database"
	"Gymnasium_reservation_WeChat_mini_program/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteAnn(c *gin.Context) {
	db := database.InitGormDB()
	id := c.PostForm("id")
	annID, _ := strconv.Atoi(id)
	ann := model.Announcement{
		ID: annID,
	}
	db.Delete(&ann)
	c.JSON(http.StatusOK, gin.H{
		"msg": "成功删除！",
	})
	defer db.Close()

}
