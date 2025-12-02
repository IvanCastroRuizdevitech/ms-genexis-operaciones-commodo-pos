package handler_sales

import (
	"net/http"
	"strconv"

	container_sales "ms-genexis-pos-operaciones/context/sales/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"

	"github.com/gin-gonic/gin"
)

func GetTransactionsByDispenserAndFaceHandler(ctx *gin.Context) {
	dispenserStr := ctx.Param("dispenserId")
	faceStr := ctx.Param("face")

	dispenserID, err := strconv.Atoi(dispenserStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, entities_main.NewErrorResponse[interface{}]("dispenserId inválido", err))
		return
	}
	face, err := strconv.Atoi(faceStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, entities_main.NewErrorResponse[interface{}]("face inválido", err))
		return
	}

	response, err := container_sales.ResolveGetTransactionsByDispenserAndFaceContainer().Execute(dispenserID, face)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error interno", err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
