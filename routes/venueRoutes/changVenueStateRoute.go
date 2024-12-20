package venueRoutes

import (
	"Gymnasium_reservation_WeChat_mini_program/controller/venueCtrl/checkAndChangeVenueState"
	"github.com/gin-gonic/gin"
)

func ChangeVenueStateRoute(r *gin.Engine) *gin.Engine {
	r.POST("/venueInfo/venueState/checkAndChangeVenueState", checkAndChangeVenueState.CheckAndChangeVenueState)
	return r
}
