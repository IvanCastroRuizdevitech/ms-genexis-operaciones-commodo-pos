package handler_creditcustomers

import (
	"net/http"

	entities_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	container_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"

	"github.com/gin-gonic/gin"
)

func InsertPreAuthorizationCustomerHandler(ctx *gin.Context) {
	var request entities_creditcustomers.InsertPreAuthorizationCustomerRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, entities_main.NewErrorResponse[interface{}]("Body invalido", err))
		return
	}

	response, err := container_creditcustomers.ResolveInsertPreAuthorizationCustomerContainer().Execute(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error interno", err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
