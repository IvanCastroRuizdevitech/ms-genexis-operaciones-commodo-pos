package handler_sales

import (
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    container_sales "ms-genexis-pos-operaciones/context/sales/presentation/container"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
    "net/http"

    "github.com/gin-gonic/gin"
)

func CheckDatafonoCancellationsInProgressHandler(ctx *gin.Context) {
    rawBody := ctx.MustGet("validatedBody")
    body := rawBody.(entities_sales.DatafonoCancellationsInProgressRequest)

    response, err := container_sales.ResolveDatafonoCancellationsInProgressContainer().Execute(&body)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error interno", err))
        return
    }

    ctx.JSON(http.StatusOK, response)
}

