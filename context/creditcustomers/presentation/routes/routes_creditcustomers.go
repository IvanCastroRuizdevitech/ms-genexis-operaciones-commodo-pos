package routes_creditcustomers

import (
	"ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	handler_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/presentation/handler_creditcustomers"
	presentation_api_middlewares "ms-genexis-pos-operaciones/presentation/api/middlewares"

	"github.com/gin-gonic/gin"
)

func LoadCreditCustomersRoutes(router *gin.RouterGroup) {
	creditCustomersGroup := router.Group("/credit-customers")
	{
		creditCustomersGroup.GET(
			"",
			handler_creditcustomers.GetCreditCustomersHandler,
		)

		creditCustomersGroup.PATCH(
			"/:id/limit",
			presentation_api_middlewares.ValidateBodyStruct[entities.UpdateCreditCustomerLimitRequest](),
			handler_creditcustomers.UpdateCreditCustomerLimitHandler,
		)
	}
}
