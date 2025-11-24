package handler_users

import (
	"net/http"

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

	response, err := container_users.ResolveAssignTagTransmissionsContainer().Execute(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Internal error", err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
