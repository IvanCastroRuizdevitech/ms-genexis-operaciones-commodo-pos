package routes_users

import (
	handler_users "ms-genexis-pos-operaciones/context/users/presentation/handler_users"

	"github.com/gin-gonic/gin"
)

func LoadUsersRoutes(router *gin.RouterGroup) {
	usersGroup := router.Group("/users")
	{
		usersGroup.GET("/getusersfilters", handler_users.GetUsersHandler)
	}
}
