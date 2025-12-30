package handler_configuration

import (
	"log"
	"net/http"
	"time"

	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	container_configuration "ms-genexis-pos-operaciones/context/configuration/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"

	"github.com/gin-gonic/gin"
)

func GetSynchronizationHandler(ctx *gin.Context) {
	raw, exists := ctx.Get("SynchronizationQuery")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Parametros de consulta no provistos"})
		return
	}

	params := raw.(entities.SynchronizationQuery)
	if _, err := time.Parse("2006-01-02", params.FechaInicio); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "fecha_inicio debe tener formato YYYY-MM-DD"})
		return
	}
	if _, err := time.Parse("2006-01-02", params.FechaFin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "fecha_fin debe tener formato YYYY-MM-DD"})
		return
	}

	fechaInicio := params.FechaInicio + " 00:00:00"
	fechaFin := params.FechaFin + " 23:59:59"

	response, err := container_configuration.ResolveGetSynchronizationContainer().Execute(
		params.IDSincronizacion,
		fechaInicio,
		fechaFin,
	)
	if err != nil {
		log.Printf("[GetSynchronizationHandler] error: %v", err)
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error interno cargando sincronizacion", err))
		return
	}

	if response == nil {
		log.Printf("[GetSynchronizationHandler] empty response")
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Respuesta vacia de sincronizacion"})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
