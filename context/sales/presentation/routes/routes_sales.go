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
		salesGroup.PATCH(
			"/movements/state/:movementId",
			presentation_api_middlewares.ValidateBodyStruct[entities_sales.UpdateMovementStateRequest](),
			handler_sales.UpdateMovementStateHandler,
		)
		salesGroup.POST(
			"/assign-customer-data",
			presentation_api_middlewares.ValidateBodyStruct[entities_sales.AssignCustomerDataRequest](),
			handler_sales.AssignCustomerDataHandler,
		)
		salesGroup.PATCH(
			"/update-client-movement",
			presentation_api_middlewares.ValidateBodyStruct[entities_sales.UpdateClientMovementRequest](),
			handler_sales.UpdateClientMovementHandler,
		)
		salesGroup.GET(
			"/get-pending-sale-datafono/:id_transaccion",
			handler_sales.GetPendingSaleDatafonoHandler,
		)
		salesGroup.PATCH(
			"/update-payment-methods",
			presentation_api_middlewares.ValidateBodyStruct[entities_sales.UpdatePaymentMethodsRequest](),
			handler_sales.UpdatePaymentMethodsHandler,
		)
		salesGroup.GET(
			"/reprint-sale/:movementId",
			handler_sales.ReprintSaleHandler,
		)
		salesGroup.GET(
			"/dispenser-details",
			handler_sales.GetDispenserDetailsHandler,
		)
		salesGroup.POST(
			"/fuel-entry-report",
			presentation_api_middlewares.ValidateBodyStruct[entities_sales.FuelEntryReportRequest](),
			handler_sales.FuelEntryReportHandler,
		)
	}
}
