package handler_creditcustomers

import (
	entities_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	container_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePreAuthorizationHandler(ctx *gin.Context) {
	rawBody := ctx.MustGet("validatedBody")
	body := rawBody.(entities_creditcustomers.PreAuthorizationRequest)

	response, err := container_creditcustomers.ResolveCreatePreAuthorizationContainer().Execute(ctx.Request.Context(), &body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("failed to create pre authorization", err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
