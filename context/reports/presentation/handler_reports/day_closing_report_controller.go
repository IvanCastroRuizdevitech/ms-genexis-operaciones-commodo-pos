package handler_reports

import (
	"log"
	container_reports "ms-genexis-pos-operaciones/context/reports/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDayClosingReportHandler(ctx *gin.Context) {
	fecha := ctx.Param("fecha")
	if fecha == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Parámetro 'fecha' es requerido"})
		return
	}

	response, err := container_reports.ResolveDayClosingReportContainer().Execute(fecha)
	log.Printf("Respuesta del GetDayClosingReportHandler: %+v \n\n", response)
	if err != nil {
		log.Printf("[GetDayClosingReportHandler] fecha=%s error: %v", fecha, err)
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error interno obteniendo cierre de dia", err))
		return
	}

	if response == nil {
		log.Printf("[GetDayClosingReportHandler] fecha=%s empty response", fecha)
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Respuesta vacia del reporte de cierre de dia"})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
