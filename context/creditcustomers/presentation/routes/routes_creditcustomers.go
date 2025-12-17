package routes_creditcustomers

import (
	handler_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/presentation/handler_creditcustomers"

	"github.com/gin-gonic/gin"
)

func LoadCreditCustomersRoutes(router *gin.RouterGroup) {
	creditCustomersGroup := router.Group("/credit-customers")
	{
		creditCustomersGroup.GET(
			"/dispenser-details",
			handler_creditcustomers.GetDispenserDetailsHandler,
		)
		creditCustomersGroup.POST(
			"/dispenser-details/families",
			handler_creditcustomers.GetDispenserDetailsByFamiliesHandler,
		)
		creditCustomersGroup.GET(
			"/identifier-types",
			handler_creditcustomers.GetIdentifierTypesHandler,
		)
		creditCustomersGroup.GET(
			"/price-families",
			handler_creditcustomers.GetPriceFamiliesHandler,
		)
		creditCustomersGroup.PATCH(
			"/transactions/by-authorization",
			handler_creditcustomers.UpdateTransactionByAuthorizationHandler,
		)
	}
}
