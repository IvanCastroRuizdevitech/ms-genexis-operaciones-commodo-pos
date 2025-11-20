package handler_home

import (
	"net/http"
	"strconv"

	container_home "ms-genexis-pos-operaciones/context/home/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"

	"github.com/gin-gonic/gin"
)

func GetMunicipalityLocationHandler(ctx *gin.Context) {
	idParam := ctx.Param("id")
	if idParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Parameter 'id' is required"})
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Parameter 'id' must be a positive integer"})
		return
	}

	response, err := container_home.ResolveMunicipalityLocationContainer().Execute(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Internal error", err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
