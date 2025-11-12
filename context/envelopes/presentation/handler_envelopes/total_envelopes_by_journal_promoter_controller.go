package handler_enveloper

import (
	"ms-genexis-pos-operaciones/context/envelopes/domain/entities"
	container_envelopes "ms-genexis-pos-operaciones/context/envelopes/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TotalEnvelopesByJournalPromoterHandler(ctx *gin.Context) {
	rawBody := ctx.MustGet("validatedBody")
	body := rawBody.(entities.EnvelopesTotalRequest)

	response, err := container_envelopes.ResolveEnvelopesContainer().Execute(&body)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error interno", err))
		return
	}

	ctx.JSON(http.StatusOK, response)

}
