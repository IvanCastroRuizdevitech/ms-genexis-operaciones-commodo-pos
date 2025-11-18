package handler_reports

import (
	"log"
	"net/http"

	entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
	container_reports "ms-genexis-pos-operaciones/context/reports/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"

	"github.com/gin-gonic/gin"
)

func GetDailyNoveltiesHandler(ctx *gin.Context) {
	var request entities_reports.DailyNoveltiesRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid request body", "error": err.Error()})
		return
	}

	if request.Ano == 0 || request.Mes == 0 || request.Dia == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Fields 'ano', 'mes' and 'dia' are required"})
		return
	}

	log.Printf("GetDailyNoveltiesHandler request: %+v\n", request)

	response, err := container_reports.ResolveDailyNoveltiesContainer().Execute(request.Ano, request.Mes, request.Dia)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Unable to obtain novelties", err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
