package handler_users

import (
	"log"
	container_users "ms-genexis-pos-operaciones/context/users/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsersHandler(ctx *gin.Context) {
	response, err := container_users.ResolveUsersContainer().Execute()
	if err != nil {
		log.Printf("[GetUsersHandler] error fetching users: %v", err)
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Internal error", err))
		return
	}

	if response == nil {
		log.Printf("[GetUsersHandler] empty response")
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Empty response fetching users", nil))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
