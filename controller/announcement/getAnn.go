package announcement

import (
	"Gymnasium_reservation_WeChat_mini_program/common/database"
	"Gymnasium_reservation_WeChat_mini_program/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAnn(c *gin.Context) {
	db := database.InitGormDB()
	var annList []model.Announcement
	db.Find(&annList)
	c.JSON(http.StatusOK, annList)
	defer db.Close()
}
