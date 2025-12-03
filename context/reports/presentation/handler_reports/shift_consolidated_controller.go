package handler_reports

import (
	"net/http"
	"time"

	entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
	container_reports "ms-genexis-pos-operaciones/context/reports/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"

	"github.com/gin-gonic/gin"
)

func GetShiftConsolidatedHandler(ctx *gin.Context) {
	var request entities_reports.ShiftConsolidatedRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, entities_main.NewErrorResponse[interface{}]("Body inválido", err))
		return
	}

	if request.FechaInicio == "" || request.FechaFin == "" {
		ctx.JSON(http.StatusBadRequest, entities_main.NewErrorResponse[interface{}]("fecha_inicio y fecha_fin son requeridos", nil))
		return
	}

	layout := "2006-01-02 15:04:05"
	if _, err := time.Parse(layout, request.FechaInicio); err != nil {
		ctx.JSON(http.StatusBadRequest, entities_main.NewErrorResponse[interface{}]("fecha_inicio debe tener formato YYYY-MM-DD HH:MM:SS", err))
		return
	}
	if _, err := time.Parse(layout, request.FechaFin); err != nil {
		ctx.JSON(http.StatusBadRequest, entities_main.NewErrorResponse[interface{}]("fecha_fin debe tener formato YYYY-MM-DD HH:MM:SS", err))
		return
	}

	response, err := container_reports.ResolveGetShiftConsolidatedContainer().Execute(
		request.FechaInicio,
		request.FechaFin,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error interno", err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
