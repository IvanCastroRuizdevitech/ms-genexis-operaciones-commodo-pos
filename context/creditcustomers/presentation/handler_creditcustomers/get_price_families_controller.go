package handler_creditcustomers

import (
	"net/http"

	container_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"

	"github.com/gin-gonic/gin"
)

func GetPriceFamiliesHandler(ctx *gin.Context) {
	response, err := container_creditcustomers.ResolveGetPriceFamiliesContainer().Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error interno", err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
