package routes_comanda

import (
	entities_comanda "ms-genexis-pos-operaciones/context/comanda/domain/entities"
	handler_comanda "ms-genexis-pos-operaciones/context/comanda/presentation/handler_comanda"
	presentation_api_middlewares "ms-genexis-pos-operaciones/presentation/api/middlewares"

	"github.com/gin-gonic/gin"
)

func LoadComandaRoutes(router *gin.RouterGroup) {
	comandaGroup := router.Group("/comanda")
	{
		comandaGroup.POST(
			"/update-comanda",
			presentation_api_middlewares.ValidateBodyStruct[entities_comanda.UpdateComandaStatusRequest](),
			handler_comanda.UpdateComandaStatusHandler,
		)
	}
}
