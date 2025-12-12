package routes_configuration

import (
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	handler_configuration "ms-genexis-pos-operaciones/context/configuration/presentation/handler_configuration"
	presentation_api_middlewares "ms-genexis-pos-operaciones/presentation/api/middlewares"

	"github.com/gin-gonic/gin"
)

func LoadConfigurationRoutes(router *gin.RouterGroup) {
	configurationGroup := router.Group("/configuration")
	{
		configurationGroup.GET(
			"/parameters",
			handler_configuration.GetParametersHandler,
		)
		configurationGroup.GET(
			"/promoter-duty",
			handler_configuration.GetPromoterDutyHandler,
		)
		configurationGroup.GET(
			"/configuracion-inicial",
			handler_configuration.GetInitialConfigurationHandler,
		)
		configurationGroup.GET(
			"/printer-ip",
			handler_configuration.GetPrinterIPHandler,
		)
		configurationGroup.GET(
			"/consecutivos-unificados",
			handler_configuration.GetUnifiedConsecutivesHandler,
		)
		configurationGroup.PUT(
			"/printer-ip",
			handler_configuration.UpdatePrinterIPHandler,
		)
		configurationGroup.POST(
			"/person-validation-admin",
			presentation_api_middlewares.ValidateBodyStruct[entities.AdminValidationRequest](),
			handler_configuration.ValidateAdminPersonHandler,
		)
	}
}
