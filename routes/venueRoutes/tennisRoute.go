package venueRoutes

import (
	"Gymnasium_reservation_WeChat_mini_program/controller/venue/tennis"
	"github.com/gin-gonic/gin"
)

func TennisRoute(r *gin.Engine) *gin.Engine {
	r.GET("/venueInfo/venueState/tennis/getVenueState/today", tennis.GetTennisVenueState1)
	r.GET("/venueInfo/venueState/tennis/getVenueState/tomorrow", tennis.GetTennisVenueState2)
	r.POST("/venueInfo/venueState/tennis/checkAndChangeVenueState/today", tennis.CheckAndChangeTennisVenueState1)
	r.POST("/venueInfo/venueState/tennis/checkAndChangeVenueState/tomorrow", tennis.CheckAndChangeTennisVenueState2)

	return r
}
