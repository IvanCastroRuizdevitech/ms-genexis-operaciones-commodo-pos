package handler_sales

import (
	"net/http"
	"strconv"

	entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
	container_sales "ms-genexis-pos-operaciones/context/sales/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"

	"github.com/gin-gonic/gin"
)

func UpdateMovementStateHandler(ctx *gin.Context) {
	movementIdStr := ctx.Param("movementId")
	movementId, err := strconv.Atoi(movementIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, entities_main.NewErrorResponse[interface{}]("id_movimiento inválido", err))
		return
	}

	rawBody := ctx.MustGet("validatedBody")
	body := rawBody.(entities_sales.UpdateMovementStateRequest)

	response, err := container_sales.ResolveUpdateMovementStateContainer().Execute(movementId, &body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error interno", err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
