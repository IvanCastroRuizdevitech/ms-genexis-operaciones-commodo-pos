package handler_reports

import (
	"net/http"

	container_reports "ms-genexis-pos-operaciones/context/reports/presentation/container"

	"github.com/gin-gonic/gin"
)

func GetTanksHandler(ctx *gin.Context) {
	response, err := container_reports.ResolveTanksContainer().Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to obtain tanks", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
