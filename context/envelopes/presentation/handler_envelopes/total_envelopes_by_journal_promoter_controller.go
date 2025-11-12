package handler_enveloper

import (
	"ms-genexis-pos-operaciones/context/envelopes/domain/entities"
	container_envelopes "ms-genexis-pos-operaciones/context/envelopes/presentation/container"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func TotalEnvelopesByJournalPromoterHandler(ctx *gin.Context) {
	rawBody := ctx.MustGet("validatedBody")
	body := rawBody.(entities.EnvelopesTotalRequest)

	response, err := container_envelopes.ResolveEnvelopesContainer().Execute(&body)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsgs(err, http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func errorMsgs(err error, code int) *entities.ResponseEvelopesTotal {
	errorJson := &entities.ResponseEvelopesTotal{
		Status:      code,
		ProcessDate: time.Now().Format("2006-01-02T15:04:05Z07:00"),
		Message:     err.Error(),
	}

	return errorJson
}
