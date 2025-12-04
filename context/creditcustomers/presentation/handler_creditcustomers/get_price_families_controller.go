package handler_creditcustomers

import (
	container_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPriceFamiliesHandler(ctx *gin.Context) {
	response, err := container_creditcustomers.ResolveGetPriceFamiliesContainer().Execute(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("failed to get price families", err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
