package handler_shift

import (
    container_shift "ms-genexis-pos-operaciones/context/shift/presentation/container"
    "ms-genexis-pos-operaciones/context/shift/domain/entities"
    "net/http"
)

import "github.com/gin-gonic/gin"

func GetDailyIncomeMeasurementsHandler(ctx *gin.Context) {
    raw, exists := ctx.Get("DailyIncomeMeasurementsQuery")
    if !exists {
        ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Parámetros de consulta no provistos"})
        return
    }

    params := raw.(entities.DailyIncomeMeasurementsQuery)

    response, err := container_shift.ResolveDailyIncomeMeasurementsContainer().ExecuteDailyIncomeMeasurements(&params)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, errorMsgs(err, http.StatusInternalServerError))
        return
    }

    ctx.JSON(http.StatusOK, response)
}

