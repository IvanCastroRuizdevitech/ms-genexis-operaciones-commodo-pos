package handler_home

import (
	container_home "ms-genexis-pos-operaciones/context/home/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadErrorNotificationHandler(ctx *gin.Context) {
	response, err := container_home.ResolveLoadErrorNotificationContainer().Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error interno", err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
