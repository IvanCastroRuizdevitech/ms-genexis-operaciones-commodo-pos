package handler_shift

import (
	"ms-genexis-pos-operaciones/context/shift/domain/entities"
	container_shift "ms-genexis-pos-operaciones/context/shift/presentation/container"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func OpeningShiftHandler(ctx *gin.Context) {
	rawBody := ctx.MustGet("validatedBody")
	body := rawBody.(entities.OpeningShiftRequest)

	response, err := container_shift.ResolveOpeningShiftContainer().ExecuteOpeningShift(&body)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsgs(err, http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func errorMsgs(err error, code int) *entities.ResponseShift {
	errorJson := &entities.ResponseShift{
		Status:      code,
		ProcessDate: time.Now().Format("2006-01-02T15:04:05Z07:00"),
		Message:     err.Error(),
	}

	return errorJson
}
