package routes_sales

import (
	entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
	handler_sales "ms-genexis-pos-operaciones/context/sales/presentation/handler_sales"
	presentation_api_middlewares "ms-genexis-pos-operaciones/presentation/api/middlewares"

	"github.com/gin-gonic/gin"
)

func LoadSalesRoutes(router *gin.RouterGroup) {
	salesGroup := router.Group("/sales")
	{
		salesGroup.POST(
			"/check-pending-sales",
			presentation_api_middlewares.ValidateBodyStruct[entities_sales.CheckPendingSalesRequest](),
			handler_sales.CheckPendingSalesHandler,
		)
		salesGroup.POST(
			"/check-ready-sales",
			presentation_api_middlewares.ValidateBodyStruct[entities_sales.CheckReadySalesRequest](),
			handler_sales.CheckReadySalesHandler,
		)
		salesGroup.POST(
			"/datafono-cancellations-in-progress",
			presentation_api_middlewares.ValidateBodyStruct[entities_sales.DatafonoCancellationsInProgressRequest](),
			handler_sales.CheckDatafonoCancellationsInProgressHandler,
		)
		salesGroup.GET(
			"/unresolved/attributes/:movementId",
			handler_sales.GetUnresolvedSaleAttributesHandler,
		)
	}
}
