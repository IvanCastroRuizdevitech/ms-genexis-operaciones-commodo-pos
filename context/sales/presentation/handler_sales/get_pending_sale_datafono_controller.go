package handler_sales

import (
    "net/http"
    "strconv"

    container_sales "ms-genexis-pos-operaciones/context/sales/presentation/container"
    entities_main "ms-genexis-pos-operaciones/domain/entities"

    "github.com/gin-gonic/gin"
)

func GetPendingSaleDatafonoHandler(ctx *gin.Context) {
    idTransaccionStr := ctx.Param("id_transaccion")
    idTransaccion, err := strconv.Atoi(idTransaccionStr)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, entities_main.NewErrorResponse[interface{}]("id_transaccion inválido", err))
        return
    }

    response, err := container_sales.ResolveGetPendingSaleDatafonoContainer().Execute(idTransaccion)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error interno", err))
        return
    }

    ctx.JSON(http.StatusOK, response)
}

