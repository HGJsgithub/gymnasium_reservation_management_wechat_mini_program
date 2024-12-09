package venueRoutes

import (
	"Gymnasium_reservation_WeChat_mini_program/controller/venue/tableTennis"
	"github.com/gin-gonic/gin"
)

func TableTennisRoute(r *gin.Engine) *gin.Engine {
	r.GET("/api/venueInfo/venueState/tableTennis/getVenueState/today", tableTennis.GetTableTennisVenueState1)
	r.GET("/api/venueInfo/venueState/tableTennis/getVenueState/tomorrow", tableTennis.GetTableTennisVenueState2)
	r.POST("/api/venueInfo/venueState/tableTennis/checkAndChangeVenueState/today", tableTennis.CheckAndChangeTableTennisVenueState1)
	r.POST("/api/venueInfo/venueState/tableTennis/checkAndChangeVenueState/tomorrow", tableTennis.CheckAndChangeTableTennisVenueState2)
	return r
}
