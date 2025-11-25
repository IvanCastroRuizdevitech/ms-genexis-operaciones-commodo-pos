package handler_reports

import (
	"log"
	"net/http"

	container_reports "ms-genexis-pos-operaciones/context/reports/presentation/container"

	"github.com/gin-gonic/gin"
)

func GetTanksHandler(ctx *gin.Context) {
	response, err := container_reports.ResolveTanksContainer().Execute()
	if err != nil {
		log.Printf("[GetTanksHandler] error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to obtain tanks", "error": err.Error()})
		return
	}

	if response == nil {
		log.Printf("[GetTanksHandler] empty response")
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Empty response for tanks"})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
