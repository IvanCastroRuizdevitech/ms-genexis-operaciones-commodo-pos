package handler_comanda

import (
	"net/http"

	entities_comanda "ms-genexis-pos-operaciones/context/comanda/domain/entities"
	container_comanda "ms-genexis-pos-operaciones/context/comanda/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"

	"github.com/gin-gonic/gin"
)

func UpdateComandaStatusHandler(ctx *gin.Context) {
	rawBody := ctx.MustGet("validatedBody")
	body := rawBody.(entities_comanda.UpdateComandaStatusRequest)

	response, err := container_comanda.ResolveUpdateComandaStatusContainer().Execute(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error interno", err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
