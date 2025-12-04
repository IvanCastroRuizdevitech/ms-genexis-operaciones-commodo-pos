package handler_sales

import (
	"net/http"
	"strconv"

	entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
	container_sales "ms-genexis-pos-operaciones/context/sales/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"

	"github.com/gin-gonic/gin"
)

func SetInvoiceAttributesHandler(ctx *gin.Context) {
	caraParam := ctx.Param("cara")
	cara, err := strconv.Atoi(caraParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, entities_main.NewErrorResponse[interface{}]("cara debe ser entero", err))
		return
	}

	req := entities_sales.SetInvoiceAttributesRequest{
		Cara: cara,
	}

	response, err := container_sales.ResolveSetInvoiceAttributesContainer().Execute(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error al actualizar atributos de factura", err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
