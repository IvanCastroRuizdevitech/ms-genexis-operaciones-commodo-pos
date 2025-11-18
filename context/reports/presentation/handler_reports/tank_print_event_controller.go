package handler_reports

import (
	"net/http"

	entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
	container_reports "ms-genexis-pos-operaciones/context/reports/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"

	"github.com/gin-gonic/gin"
)

func CreateTankPrintEventHandler(ctx *gin.Context) {
	var request entities_reports.TankPrintEventRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid request body", "error": err.Error()})
		return
	}

	if len(request.TankIDs) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Field 'tank_ids' is required and must contain at least one tank id"})
		return
	}

	response, err := container_reports.ResolveTankPrintEventContainer().Execute(request.TankIDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Failed to create tank print event.", err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
