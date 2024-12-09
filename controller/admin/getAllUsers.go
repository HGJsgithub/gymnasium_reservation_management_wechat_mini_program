package admin

import (
	"Gymnasium_reservation_WeChat_mini_program/common/database"
	"Gymnasium_reservation_WeChat_mini_program/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func GetAllUsers(c *gin.Context) {
	db := database.ConnectTOGormDB()
	var userList []model.User
	db.Find(&userList)
	c.JSON(http.StatusOK, userList)
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)

}
