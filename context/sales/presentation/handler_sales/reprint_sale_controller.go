package handler_sales

import (
    "net/http"
    "strconv"

    container_sales "ms-genexis-pos-operaciones/context/sales/presentation/container"
    entities_main "ms-genexis-pos-operaciones/domain/entities"

    "github.com/gin-gonic/gin"
)

func ReprintSaleHandler(ctx *gin.Context) {
    movementIdStr := ctx.Param("movementId")
    movementId, err := strconv.Atoi(movementIdStr)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, entities_main.NewErrorResponse[interface{}]("movementId inválido", err))
        return
    }

    response, err := container_sales.ResolveReprintSaleContainer().Execute(movementId)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error interno", err))
        return
    }

    ctx.JSON(http.StatusOK, response)
}

