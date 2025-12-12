package handler_configuration

import (
	"net/http"

	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	container_configuration "ms-genexis-pos-operaciones/context/configuration/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"

	"github.com/gin-gonic/gin"
)

func ValidateAdminPersonHandler(ctx *gin.Context) {
	rawBody := ctx.MustGet("validatedBody")
	body := rawBody.(entities.AdminValidationRequest)

	response, err := container_configuration.ResolveValidateAdminPersonContainer().Execute(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error interno", err))
		return
	}

	payload := gin.H{
		"status":        response.Status,
		"success":       response.Success,
		"authenticated": false,
		"message":       response.Message,
	}

	if response.Data != nil {
		payload["success"] = response.Data.Success
		payload["authenticated"] = response.Data.Authenticated
		payload["message"] = response.Data.Message
	}

	ctx.JSON(http.StatusOK, payload)
}
