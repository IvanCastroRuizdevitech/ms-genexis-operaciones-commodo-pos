package handler_sales

import (
	"net/http"

	entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
	container_sales "ms-genexis-pos-operaciones/context/sales/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"

	"github.com/gin-gonic/gin"
)

func UpsertActiveDispenserFaceTransactionHandler(ctx *gin.Context) {
	rawBody := ctx.MustGet("validatedBody")
	body := rawBody.(entities_sales.ActiveDispenserFaceTransactionRequest)

	response, err := container_sales.ResolveUpsertActiveDispenserFaceTransactionContainer().Execute(
		body.Surtidor,
		body.Cara,
		body.Codigo,
		body.Grado,
		body.ProveedorID,
		body.MontoMaximo,
		body.CantidadMaxima,
		body.Trama,
		body.PromotorID,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error interno", err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
