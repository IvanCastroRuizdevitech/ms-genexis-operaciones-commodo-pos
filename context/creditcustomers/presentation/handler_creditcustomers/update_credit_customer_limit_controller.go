package handler_creditcustomers

import (
	"net/http"
	"strconv"

	"ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	container_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"

	"github.com/gin-gonic/gin"
)

func UpdateCreditCustomerLimitHandler(ctx *gin.Context) {
	rawBody := ctx.MustGet("validatedBody")
	body := rawBody.(entities.UpdateCreditCustomerLimitRequest)

	if idParam := ctx.Param("id"); idParam != "" {
		if id, err := strconv.ParseInt(idParam, 10, 64); err == nil {
			body.CustomerID = id
		} else {
			ctx.JSON(http.StatusBadRequest, entities_main.NewErrorResponse[interface{}]("invalid customer id", err))
			return
		}
	}

	response, err := container_creditcustomers.ResolveUpdateCreditCustomerLimitContainer().Execute(ctx.Request.Context(), &body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("failed to update credit customer limit", err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
