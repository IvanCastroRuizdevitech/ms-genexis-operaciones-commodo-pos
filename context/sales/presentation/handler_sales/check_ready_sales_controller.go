package handler_sales

import (
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    container_sales "ms-genexis-pos-operaciones/context/sales/presentation/container"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
    "net/http"

    "github.com/gin-gonic/gin"
)

func CheckReadySalesHandler(ctx *gin.Context) {
    rawBody := ctx.MustGet("validatedBody")
    body := rawBody.(entities_sales.CheckReadySalesRequest)

    response, err := container_sales.ResolveReadySalesContainer().Execute(&body)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error interno", err))
        return
    }

    ctx.JSON(http.StatusOK, response)
}

