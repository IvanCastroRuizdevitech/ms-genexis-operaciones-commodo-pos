package handler_shift

import (
	container_shift "ms-genexis-pos-operaciones/context/shift/presentation/container"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FuelPumpsHandler(ctx *gin.Context) {
	response, err := container_shift.ResolveFuelPumpsContainer().ExecuteFuelPumps()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsgs(err, http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
