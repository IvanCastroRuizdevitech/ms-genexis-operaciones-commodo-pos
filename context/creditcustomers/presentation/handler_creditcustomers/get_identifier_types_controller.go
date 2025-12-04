package handler_creditcustomers

import (
	container_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIdentifierTypesHandler(ctx *gin.Context) {
	response, err := container_creditcustomers.ResolveGetIdentifierTypesContainer().Execute(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("failed to get identifier types", err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
