package api_routes

import (
	"time"

	routes_comanda "ms-genexis-pos-operaciones/context/comanda/presentation/routes"
	routes_configuration "ms-genexis-pos-operaciones/context/configuration/presentation/routes"
	"ms-genexis-pos-operaciones/domain/constants"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

// enableSwaggerDocs controla la exposición de rutas de documentación Swagger.
// Debe permanecer en true para ambientes de desarrollo y deshabilitarse para PR/producción.
const enableSwaggerDocs = true

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
	// Documentación de APIs por contexto usando Swagger (OpenAPI)
	// Expone (solo si enableSwaggerDocs == true):
	//  - GET /docs              -> UI de Swagger (via CDN)
	//  - GET /docs/swagger.json -> Documento OpenAPI agrupado por contexto

	api := router.Group(constants.API_PATH)

	routes_comanda.LoadComandaRoutes(api)
	routes_configuration.LoadConfigurationRoutes(api)

	return router, nil
}
