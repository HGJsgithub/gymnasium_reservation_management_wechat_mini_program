package venueRoutes

import (
	"Gymnasium_reservation_WeChat_mini_program/controller/venue/badminton"
	"github.com/gin-gonic/gin"
)

func BadmintonRoute(r *gin.Engine) *gin.Engine {
	//获取今天的场地状态
	r.GET("/api/venueInfo/venueState/badminton/getVenueState/today", badminton.GetBadmintonVenueState1)
	//获取明天的场地状态
	r.GET("/api/venueInfo/venueState/badminton/getVenueState/tomorrow", badminton.GetBadmintonVenueState2)
	//改变今天的场地状态
	r.POST("/api/venueInfo/venueState/badminton/checkAndChangeVenueState/today", badminton.CheckAndChangeBadmintonVenueState1)
	//改变明天的场地状态
	r.POST("/api/venueInfo/venueState/badminton/checkAndChangeVenueState/tomorrow", badminton.CheckAndChangeBadmintonVenueState2)
	return r
}
