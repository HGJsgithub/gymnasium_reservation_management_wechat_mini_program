package venueRoutes

import (
	"github.com/gin-gonic/gin"
)

func VenueRoutes(r *gin.Engine) *gin.Engine {
	r = BadmintonRoute(r)
	r = TableTennisRoute(r)
	r = TennisRoute(r)

	return r
}
