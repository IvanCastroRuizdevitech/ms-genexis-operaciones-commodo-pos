package handler_configuration

import (
	"net/http"

	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	container_configuration "ms-genexis-pos-operaciones/context/configuration/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"

	"github.com/gin-gonic/gin"
)

func UpdatePrinterIPHandler(ctx *gin.Context) {
	body := entities.PrinterIPUpdateRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, entities_main.NewErrorResponse[interface{}]("Payload invalido", err))
		return
	}

	response, err := container_configuration.ResolveUpdatePrinterIPContainer().Execute(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error interno", err))
		return
	}

	dataPayload := map[string]any{}
	mensaje := ""
	if response.Data != nil {
		if response.Data.Formatted != nil {
			dataPayload = response.Data.Formatted
		}
		if response.Data.Message != "" {
			mensaje = response.Data.Message
		} else if response.Data.Raw != "" {
			mensaje = response.Data.Raw
		}
	}
	if mensaje == "" {
		mensaje = response.Message
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  response.Status,
		"success": response.Success,
		"data":    dataPayload,
		"mensaje": mensaje,
	})
}
