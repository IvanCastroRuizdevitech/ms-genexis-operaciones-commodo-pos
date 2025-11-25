package handler_users

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	"ms-genexis-pos-operaciones/context/users/domain/entities"
	container_users "ms-genexis-pos-operaciones/context/users/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

func AssignTagTransmissionsHandler(ctx *gin.Context) {
	var body entities.AssignTagTransmissionsRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, entities_main.NewErrorResponse[interface{}]("Invalid request body", err))
		return
	}

	if body.Tag == "" || body.Identification == "" || body.Medio == "" {
		ctx.JSON(http.StatusBadRequest, entities_main.NewErrorResponse[interface{}]("Fields 'tag', 'identification' and 'medio' are required", nil))
		return
	}

	// Paso 1: asignar/actualizar tag (limpia tag previo y asigna a la identificación)
	assignResp, err := container_users.ResolveAssignTagContainer().Execute(&entities.AssignTagRequest{
		Tag:            body.Tag,
		Identification: body.Identification,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, entities_main.NewErrorResponse[interface{}]("User not found for identification", err))
			return
		}
		log.Printf("[AssignTagTransmissionsHandler] error assigning tag: %v", err)
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Internal error", err))
		return
	}

	// Paso 2: generar transmisiones (si falla, devolvemos 500)
	transResp, err := container_users.ResolveAssignTagTransmissionsContainer().Execute(&body)
	if err != nil {
		log.Printf("[AssignTagTransmissionsHandler] error generating transmissions: %v", err)
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error generating transmissions", err))
		return
	}

	if assignResp == nil || transResp == nil {
		log.Printf("[AssignTagTransmissionsHandler] empty response assignResp=%v transResp=%v", assignResp, transResp)
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Empty response generating transmissions", nil))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":       200,
		"message":      "Tag assigned and transmissions generated",
		"process_date": transResp.ProcessDate,
		"data": gin.H{
			"id":                assignResp.Data.ID,
			"transmissions_raw": transResp.Data.TransmissionsRaw,
		},
	})
}
