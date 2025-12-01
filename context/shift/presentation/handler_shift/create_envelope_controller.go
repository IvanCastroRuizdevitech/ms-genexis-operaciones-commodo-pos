package handler_shift

import (
	"ms-genexis-pos-operaciones/context/shift/domain/entities"
	container_shift "ms-genexis-pos-operaciones/context/shift/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEnvelopeHandler(ctx *gin.Context) {
	rawBody := ctx.MustGet("validatedBody")
	body := rawBody.(entities.EnvelopeRequest)

	response, err := container_shift.ResolveCreateEnvelopeContainer().Execute(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error interno", err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
