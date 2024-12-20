package annCtrl

import (
	"Gymnasium_reservation_WeChat_mini_program/common/database"
	"Gymnasium_reservation_WeChat_mini_program/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func GetAnn(c *gin.Context) {
	db := database.ConnectToGormDB()
	var annList []model.Announcement
	db.Find(&annList)
	c.JSON(http.StatusOK, annList)
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
}
