package venueRoutes

import (
	"github.com/gin-gonic/gin"
)

func VenueRoutes(r *gin.Engine) *gin.Engine {
	r = GetVenueStateRoutes(r)
	r = ChangeVenueStateRoute(r)
	return r
}
