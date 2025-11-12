package api_routes

import (
	routes_shift "ms-genexis-pos-operaciones/context/shift/presentation/routes"
	"ms-genexis-pos-operaciones/domain/constants"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func GinConfig() (*gin.Engine, error) {
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE, PATCH",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          300 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	api := router.Group(constants.API_PATH)
	routes_shift.LoadShiftRoutes(api)

	return router, nil
}
