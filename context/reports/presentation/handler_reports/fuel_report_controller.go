package handler_reports

import (
	"log"
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
		log.Printf("[GetFuelReportHandler] fecha=%s error: %v", fecha, err)
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error interno obteniendo reporte de combustible", err))
		return
	}

	if response == nil {
		log.Printf("[GetFuelReportHandler] fecha=%s empty response", fecha)
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Respuesta vacia del reporte de combustible"})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
