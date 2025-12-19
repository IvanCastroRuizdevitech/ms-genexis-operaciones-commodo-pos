package handler_sales

import (
	"net/http"

	container_sales "ms-genexis-pos-operaciones/context/sales/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"

	"github.com/gin-gonic/gin"
)

func GetVehicleTypesHandler(ctx *gin.Context) {
	response, err := container_sales.ResolveGetVehicleTypesContainer().Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error interno", err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
