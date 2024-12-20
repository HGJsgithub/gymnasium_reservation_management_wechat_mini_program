package venueRoutes

import (
	"Gymnasium_reservation_WeChat_mini_program/controller/venueCtrl/getVenueState"
	"github.com/gin-gonic/gin"
)

func GetVenueStateRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/venueInfo/venueState/getVenueState/badminton", getVenueState.GetBadmintonVenueStateTableHandle)
	r.GET("/venueInfo/venueState/getVenueState/tableTennis", getVenueState.GetTableTennisVenueStateTableHandle)
	r.GET("/venueInfo/venueState/getVenueState/tennis", getVenueState.GetTennisVenueStateTableHandle)
	return r
}
