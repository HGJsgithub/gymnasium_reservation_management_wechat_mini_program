package routes

import (
	"Gymnasium_reservation_WeChat_mini_program/routes/venueRoutes"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r = AnnRoute(r)
	r = UserRoute(r)
	r = AdminRoute(r)
	r = OrderRoute(r)
	r = venueRoutes.VenueRoutes(r)
	return r
}
