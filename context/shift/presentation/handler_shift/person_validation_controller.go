package handler_shift

import (
	container_shift "ms-genexis-pos-operaciones/context/shift/presentation/container"
	"ms-genexis-pos-operaciones/context/shift/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PersonValidationHandler(ctx *gin.Context) {
	rawBody := ctx.MustGet("validatedBody")
	body := rawBody.(entities.PersonValidationRequest)

	response, err := container_shift.ResolvePersonValidationContainer().ExecutePersonValidation(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsgs(err, http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
