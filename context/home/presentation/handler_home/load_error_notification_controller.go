package handler_home

import (
	"log"
	container_home "ms-genexis-pos-operaciones/context/home/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadErrorNotificationHandler(ctx *gin.Context) {
	response, err := container_home.ResolveLoadErrorNotificationContainer().Execute()
	if err != nil {
		log.Printf("[LoadErrorNotificationHandler] error: %v", err)
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error interno cargando notificaciones", err))
		return
	}

	if response == nil {
		log.Printf("[LoadErrorNotificationHandler] empty response")
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Respuesta vacia de notificaciones"})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
