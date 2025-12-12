package handler_configuration

import (
	"net/http"

	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	container_configuration "ms-genexis-pos-operaciones/context/configuration/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"

	"github.com/gin-gonic/gin"
)

func ProcessNotificationHandler(ctx *gin.Context) {
	rawBody := ctx.MustGet("validatedBody")
	body := rawBody.(entities.ProcessNotificationRequest)

	response, err := container_configuration.ResolveProcessNotificationContainer().Execute(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error interno", err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
