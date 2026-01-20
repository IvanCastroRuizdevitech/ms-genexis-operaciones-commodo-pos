package routes_configuration

import (
	handler_configuration "ms-genexis-pos-operaciones/context/configuration/presentation/handler_configuration"

	"github.com/gin-gonic/gin"
)

func LoadConfigurationRoutes(router *gin.RouterGroup) {
	configurationGroup := router.Group("/configuration")
	{
		configurationGroup.GET(
			"/configuracion-inicial",
			handler_configuration.GetInitialConfigurationHandler,
		)
	}
}
