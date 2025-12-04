package routes_creditcustomers

import (
	handler_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/presentation/handler_creditcustomers"

	"github.com/gin-gonic/gin"
)

func LoadCreditCustomersRoutes(router *gin.RouterGroup) {
	creditCustomersGroup := router.Group("/credit-customers")
	{
		creditCustomersGroup.GET(
			"/identifier-types",
			handler_creditcustomers.GetIdentifierTypesHandler,
		)
		creditCustomersGroup.GET(
			"/price-families",
			handler_creditcustomers.GetPriceFamiliesHandler,
		)
	}
}
