package handler_users

import (
	"errors"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"ms-genexis-pos-operaciones/context/users/domain/entities"
	container_users "ms-genexis-pos-operaciones/context/users/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"

	"github.com/gin-gonic/gin"
)

func AssignTagHandler(ctx *gin.Context) {
	var body entities.AssignTagRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, entities_main.NewErrorResponse[interface{}]("Invalid request body", err))
		return
	}

	if body.Tag == "" || body.Identification == "" {
		ctx.JSON(http.StatusBadRequest, entities_main.NewErrorResponse[interface{}]("Tag and identification are required", nil))
		return
	}

	response, err := container_users.ResolveAssignTagContainer().Execute(&body)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, entities_main.NewErrorResponse[interface{}]("User not found for identification", err))
			return
		}
		log.Printf("[AssignTagHandler] error assigning tag: %v", err)
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Internal error", err))
		return
	}

	if response == nil {
		log.Printf("[AssignTagHandler] empty response")
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Empty response assigning tag", nil))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
