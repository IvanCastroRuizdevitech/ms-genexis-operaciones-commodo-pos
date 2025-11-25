package handler_reports

import (
	"log"
	"net/http"

	container_reports "ms-genexis-pos-operaciones/context/reports/presentation/container"

	"github.com/gin-gonic/gin"
)

func GetMovementTypesHandler(ctx *gin.Context) {
	response, err := container_reports.ResolveMovementTypesContainer().Execute()
	if err != nil {
		log.Printf("[GetMovementTypesHandler] error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to obtain movement types", "error": err.Error()})
		return
	}

	if response == nil {
		log.Printf("[GetMovementTypesHandler] empty response")
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Empty response for movement types"})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
