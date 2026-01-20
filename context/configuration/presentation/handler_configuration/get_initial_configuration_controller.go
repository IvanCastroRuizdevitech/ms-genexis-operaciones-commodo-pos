package handler_configuration

import (
	"net/http"

	container_configuration "ms-genexis-pos-operaciones/context/configuration/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"

	"github.com/gin-gonic/gin"
)

func GetInitialConfigurationHandler(ctx *gin.Context) {
	response, err := container_configuration.ResolveGetInitialConfigurationContainer().Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error interno", err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
