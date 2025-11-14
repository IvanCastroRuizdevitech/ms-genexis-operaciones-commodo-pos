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
    }
}

