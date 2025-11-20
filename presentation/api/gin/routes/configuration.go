package api_routes

import (
	routes_configuration "ms-genexis-pos-operaciones/context/configuration/presentation/routes"
	routes_envelopes "ms-genexis-pos-operaciones/context/envelopes/presentation/routes"
	routes_home "ms-genexis-pos-operaciones/context/home/presentation/routes"
	routes_reports "ms-genexis-pos-operaciones/context/reports/presentation/routes"
	routes_sales "ms-genexis-pos-operaciones/context/sales/presentation/routes"
	routes_shift "ms-genexis-pos-operaciones/context/shift/presentation/routes"
	"ms-genexis-pos-operaciones/domain/constants"
	"time"

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
    if enableSwaggerDocs {
        registerSwaggerRoutes(router)
    }
    api := router.Group(constants.API_PATH)
	routes_shift.LoadShiftRoutes(api)
	routes_envelopes.LoadEnvelopesRoutes(api)
	routes_configuration.LoadConfigurationRoutes(api)
	routes_reports.LoadReportsRoutes(api)
	routes_sales.LoadSalesRoutes(api)
	routes_home.LoadHomeRoutes(api)

	return router, nil
}
