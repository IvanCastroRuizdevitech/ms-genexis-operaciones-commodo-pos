package routes_shift

import (
	"github.com/gin-gonic/gin"
)

func LoadShiftRoutes(router *gin.RouterGroup) {
	_ = router.Group("/authorization/v1")
	{
		// autorizationGroup.Use(middlewares.IsBodyEmpty)
	}
}
