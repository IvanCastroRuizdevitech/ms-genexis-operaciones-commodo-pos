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
	}
}
