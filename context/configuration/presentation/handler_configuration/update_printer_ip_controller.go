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

	// Remap the response to the format required by the POS client.
	var parsed map[string]any
	if response.Data != nil && response.Data.Parsed != nil {
		parsed = response.Data.Parsed
	}

	dataPayload := map[string]any{}
	if parsed != nil {
		if dataRaw, ok := parsed["data"].(map[string]any); ok {
			if valorAnterior, ok := dataRaw["valor_anterior"]; ok {
				dataPayload["valor_anterior"] = valorAnterior
			}
			if valorNuevo, ok := dataRaw["valor_nuevo"]; ok {
				dataPayload["valor_nuevo"] = valorNuevo
			}
		}
	}

	mensaje := ""
	if parsed != nil {
		if msg, ok := parsed["mensaje"].(string); ok {
			mensaje = msg
		}
	}
	if mensaje == "" && response.Data != nil {
		mensaje = response.Data.Raw
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  response.Status,
		"success": response.Success,
		"data":    dataPayload,
		"mensaje": mensaje,
	})
}
