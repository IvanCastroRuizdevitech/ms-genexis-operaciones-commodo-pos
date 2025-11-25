package handler_home

import (
	"log"
	"net/http"
	"time"

	container_home "ms-genexis-pos-operaciones/context/home/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"

	"github.com/gin-gonic/gin"
)

func ProcessPendingTransmissionsHandler(ctx *gin.Context) {
	start := time.Now()
	defer func() {
		log.Printf("[ProcessPendingTransmissionsHandler] request duration: %s", time.Since(start))
	}()

	response, err := container_home.ResolveProcessPendingTransmissionsContainer().Execute()
	if err != nil {
		log.Printf("[ProcessPendingTransmissionsHandler] error: %v", err)
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Internal error processing pending transmissions", err))
		return
	}

	if response == nil {
		log.Printf("[ProcessPendingTransmissionsHandler] empty response from use case")
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Empty response from pending transmissions"})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
