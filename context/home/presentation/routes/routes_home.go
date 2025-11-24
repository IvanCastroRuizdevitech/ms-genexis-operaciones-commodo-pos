package routes_home

import (
	handler_home "ms-genexis-pos-operaciones/context/home/presentation/handler_home"

	"github.com/gin-gonic/gin"
)

func LoadHomeRoutes(router *gin.RouterGroup) {
	homeGroup := router.Group("/home")
	{
		homeGroup.GET("/load-error-notification", handler_home.LoadErrorNotificationHandler)
		handler_home.StartPendingTransmissionsScheduler()
		homeGroup.GET("/pending-transmissions", handler_home.ProcessPendingTransmissionsHandler)
	}
}
