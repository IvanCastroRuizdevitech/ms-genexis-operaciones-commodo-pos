package handler_reports

import (
	container_reports "ms-genexis-pos-operaciones/context/reports/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFuelReportHandler(ctx *gin.Context) {
	fecha := ctx.Param("fecha")
	if fecha == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Parametro 'fecha' es requerido"})
		return
	}

	response, err := container_reports.ResolveFuelReportContainer().Execute(fecha)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error interno", err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
