package handler_users

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
	"ms-genexis-pos-operaciones/context/users/domain/entities"
	container_users "ms-genexis-pos-operaciones/context/users/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"

	"github.com/gin-gonic/gin"
)

func AssignTagTransmissionsHandler(ctx *gin.Context) {
	var body entities.AssignTagTransmissionsRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, entities_main.NewErrorResponse[interface{}]("Invalid request body", err))
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
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Internal error", err))
		return
	}

	// Paso 2: generar transmisiones (si falla, devolvemos 500)
	transResp, err := container_users.ResolveAssignTagTransmissionsContainer().Execute(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error generating transmissions", err))
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
