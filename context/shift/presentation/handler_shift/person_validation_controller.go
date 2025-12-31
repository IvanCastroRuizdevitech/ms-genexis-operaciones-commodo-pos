package handler_shift

import (
	"net/http"
	"strconv"

	"ms-genexis-pos-operaciones/context/shift/domain/entities"
	container_shift "ms-genexis-pos-operaciones/context/shift/presentation/container"

	"github.com/gin-gonic/gin"
)

func PersonValidationHandler(ctx *gin.Context) {
	rawBody := ctx.MustGet("validatedBody")
	body := rawBody.(entities.PersonValidationRequest)

	requireAdmin := false
	requireAdminParam := ctx.Query("requiere_admin")
	if requireAdminParam == "" {
		requireAdminParam = ctx.Query("require_admin")
	}
	if requireAdminParam != "" {
		parsedRequireAdmin, err := strconv.ParseBool(requireAdminParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Parametro requiere_admin invalido"})
			return
		}
		requireAdmin = parsedRequireAdmin
	}

	response, status, err := container_shift.ResolvePersonValidationContainer().ExecutePersonValidation(&body, requireAdmin)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsgs(err, http.StatusInternalServerError))
		return
	}

	ctx.JSON(status, response)
}
