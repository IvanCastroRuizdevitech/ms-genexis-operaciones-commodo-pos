package routes_shift

import (
	"ms-genexis-pos-operaciones/context/shift/presentation/handler_shift"

	"github.com/gin-gonic/gin"
)

func LoadShiftRoutes(router *gin.RouterGroup) {
	shiftGroup := router.Group("/shift/v1")
	{
		// autorizationGroup.Use(middlewares.IsBodyEmpty)
		shiftGroup.POST("/opening", handler_shift.OpeningShiftHandler)
	}
}
