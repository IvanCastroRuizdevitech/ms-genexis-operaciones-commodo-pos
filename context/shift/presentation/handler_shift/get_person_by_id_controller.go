package handler_shift

import (
	"net/http"

	"ms-genexis-pos-operaciones/context/shift/domain/entities"
	container_shift "ms-genexis-pos-operaciones/context/shift/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"

	"github.com/gin-gonic/gin"
)

func GetPersonByIDHandler(ctx *gin.Context) {
	rawBody := ctx.MustGet("validatedBody")
	body := rawBody.(entities.PersonByIDRequest)

	response, err := container_shift.ResolveGetPersonByIDContainer().ExecutePersonByID(body.PersonaID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error interno", err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
