# Para renderizado automatico en dev
- go install github.com/air-verse/air@latest
- air init
- Iniciar el proyecto con : air

# Compilar
- GOOS=linux GOARCH=arm64 go build -o ms-genexis-operaciones-commodo-pos main.go

## Swagger
- Endpoints (solo en desarrollo):
  - `GET /docs` UI de Swagger (vía CDN)
  - `GET /docs/swagger.json` documento OpenAPI con rutas agrupadas por contexto
- Toggle: editar `enableSwaggerDocs` en `presentation/api/gin/routes/configuration.go`.
  - Desarrollo: `true` para exponer documentación.
  - PR/Producción: `false` para no exponer documentación.
