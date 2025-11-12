package routes_shift

import (
	"ms-genexis-pos-operaciones/context/shift/domain/entities"
	"ms-genexis-pos-operaciones/context/shift/presentation/handler_shift"
	presentation_api_middlewares "ms-genexis-pos-operaciones/presentation/api/middlewares"

	"github.com/gin-gonic/gin"
)

func LoadShiftRoutes(router *gin.RouterGroup) {
	shiftGroup := router.Group("/shift")
	{
		shiftGroup.POST("/opening",
			presentation_api_middlewares.ValidateBodyStruct[entities.OpeningShiftRequest](),
			handler_shift.OpeningShiftHandler,
		)
	}
}
