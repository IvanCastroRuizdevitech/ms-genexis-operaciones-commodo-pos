package handler_reports

import (
	"net/http"

	container_reports "ms-genexis-pos-operaciones/context/reports/presentation/container"

	"github.com/gin-gonic/gin"
)

func GetMovementTypesHandler(ctx *gin.Context) {
	response, err := container_reports.ResolveMovementTypesContainer().Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to obtain movement types", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
