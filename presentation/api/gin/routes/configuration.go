package api_routes

import (
	"log"
	"os"
	"strconv"
	"time"

	routes_configuration "ms-genexis-pos-operaciones/context/configuration/presentation/routes"
	routes_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/presentation/routes"
	routes_envelopes "ms-genexis-pos-operaciones/context/envelopes/presentation/routes"
	routes_home "ms-genexis-pos-operaciones/context/home/presentation/routes"
	routes_reports "ms-genexis-pos-operaciones/context/reports/presentation/routes"
	routes_sales "ms-genexis-pos-operaciones/context/sales/presentation/routes"
	routes_shift "ms-genexis-pos-operaciones/context/shift/presentation/routes"
	routes_users "ms-genexis-pos-operaciones/context/users/presentation/routes"
	"ms-genexis-pos-operaciones/domain/constants"
	handlers_printer "ms-genexis-pos-operaciones/internal/handler"
	"ms-genexis-pos-operaciones/internal/printer"

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

	// Printer micro-endpoint (POST /printer) mounted at root for local calls.
	printerTimeout := parsePrinterTimeout(os.Getenv("PRINTER_TIMEOUT_MS"))
	printerClient := &printer.Client{Timeout: printerTimeout}
	printerHandler := handlers_printer.New(
		printerClient,
		log.New(os.Stdout, "[printer] ", log.LstdFlags|log.Lmicroseconds),
	)
	printerHandler.Register(router)

	api := router.Group(constants.API_PATH)
	routes_shift.LoadShiftRoutes(api)
	routes_envelopes.LoadEnvelopesRoutes(api)
	routes_configuration.LoadConfigurationRoutes(api)
	routes_creditcustomers.LoadCreditCustomersRoutes(api)
	routes_reports.LoadReportsRoutes(api)
	routes_sales.LoadSalesRoutes(api)
	routes_home.LoadHomeRoutes(api)
	routes_users.LoadUsersRoutes(api)

	return router, nil
}

func parsePrinterTimeout(val string) time.Duration {
	if val == "" {
		return 3 * time.Second
	}
	if ms, err := strconv.Atoi(val); err == nil && ms > 0 {
		return time.Duration(ms) * time.Millisecond
	}
	return 3 * time.Second
}
