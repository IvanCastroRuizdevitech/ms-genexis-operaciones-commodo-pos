package handler_home

import (
	"log"
	"net/http"
	"time"

	container_home "ms-genexis-pos-operaciones/context/home/presentation/container"

	"github.com/gin-gonic/gin"
)

func GetPendingTransmissionsHandler(ctx *gin.Context) {
	start := time.Now()
	defer func() {
		log.Printf("[GetPendingTransmissionsHandler] request duration: %s", time.Since(start))
	}()

	response, err := container_home.ResolvePendingTransmissionsContainer().Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Internal error", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
