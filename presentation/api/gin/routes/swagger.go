package api_routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// registerSwaggerRoutes registers endpoints to serve a minimal Swagger UI and the OpenAPI spec.
func registerSwaggerRoutes(r *gin.Engine) {
	// OpenAPI JSON
	r.GET("/docs/swagger.json", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(swaggerJSON))
	})

	// Lightweight Swagger UI via CDN
	r.GET("/docs", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(swaggerHTML))
	})
}

// Minimal HTML hosting Swagger UI from CDN and pointing to our JSON.
const swaggerHTML = `<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Genexis POS Operaciones - API Docs</title>
    <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist/swagger-ui.css" />
    <style>body { margin:0; background:#fafafa; }</style>
  </head>
  <body>
    <div id="swagger-ui"></div>
    <script src="https://unpkg.com/swagger-ui-dist/swagger-ui-bundle.js"></script>
    <script>
      window.onload = function () {
        window.ui = SwaggerUIBundle({
          url: '/docs/swagger.json',
          dom_id: '#swagger-ui',
          presets: [SwaggerUIBundle.presets.apis],
          layout: 'BaseLayout'
        });
      };
    </script>
  </body>
</html>`
// Static OpenAPI 3.0 document. Paths grouped by context via tags with accurate request/response schemas.
const swaggerJSON = `{
  "openapi": "3.0.3",
  "info": {
    "title": "Genexis POS Operaciones API",
    "version": "1.0.0",
    "description": "Documentación de APIs organizada por contextos (Shift, Envelopes, Configuration, Reports, Sales)."
  },
  "servers": [{ "url": "/api/v1", "description": "Base path" }],
  "tags": [
    { "name": "Shift", "description": "Operaciones de turnos" },
    { "name": "Envelopes", "description": "Operaciones de sobres" },
    { "name": "Home", "description": "Panel principal y notificaciones" },
    { "name": "Configuration", "description": "Parámetros y configuración" },
    { "name": "Reports", "description": "Reportes" },
    { "name": "Sales", "description": "Operaciones de ventas" }
  ],
  "paths": {
    "/shift/opening": {
      "post": {
        "tags": ["Shift"],
        "summary": "Apertura de turno",
        "requestBody": {
          "required": true,
          "content": { "application/json": { "schema": { "$ref": "#/components/schemas/OpeningShiftRequest" } } }
        },
        "responses": {
          "200": { "description": "OK", "content": { "application/json": { "schema": { "$ref": "#/components/schemas/ResponseShift" } } } }
        }
      }
    },
    "/shift/daily-income-measurements": {
      "post": {
        "tags": ["Shift"],
        "summary": "Mediciones de ingresos diarios",
        "parameters": [
          { "name": "fecha_inicio", "in": "query", "required": true, "schema": { "type": "string", "format": "date-time" } },
          { "name": "fecha_fin", "in": "query", "required": true, "schema": { "type": "string", "format": "date-time" } }
        ],
        "responses": {
          "200": { "description": "OK", "content": { "application/json": { "schema": { "$ref": "#/components/schemas/ResponseShiftArray" } } } }
        }
      }
    },
    "/shift/fuel-pumps": {
      "get": {
        "tags": ["Shift"],
        "summary": "Listado de surtidores",
        "responses": {
          "200": { "description": "OK", "content": { "application/json": { "schema": { "$ref": "#/components/schemas/ResponseShiftArray" } } } }
        }
      }
    },
    "/shift/person-validation": {
      "post": {
        "tags": ["Shift"],
        "summary": "Validación de credenciales para cierre de turno",
        "requestBody": {
          "required": true,
          "content": { "application/json": { "schema": { "$ref": "#/components/schemas/PersonValidationRequest" } } }
        },
        "responses": {
          "200": { "description": "OK", "content": { "application/json": { "schema": { "$ref": "#/components/schemas/ResponsePersonShift" } } } }
        }
      }
    },
    "/envelopes/total": {
      "post": {
        "tags": ["Envelopes"],
        "summary": "Total de sobres por promotor y diario",
        "requestBody": {
          "required": true,
          "content": { "application/json": { "schema": { "$ref": "#/components/schemas/EnvelopesTotalRequest" } } }
        },
        "responses": {
          "200": { "description": "OK", "content": { "application/json": { "schema": { "$ref": "#/components/schemas/ResponseTotalEnvelopes" } } } }
        }
      }
    },
    "/home/load-error-notification": {
      "get": {
        "tags": ["Home"],
        "summary": "Obtiene errores pendientes de notificación",
        "responses": {
          "200": { "description": "OK", "content": { "application/json": { "schema": { "$ref": "#/components/schemas/ResponseErrorNotificationList" } } } }
        }
      }
    },
    "/home/pending-transmissions": {
      "get": {
        "tags": ["Home"],
        "summary": "Ejecuta la sincronización de transmisiones pendientes (tarea programada cada minuto)",
        "responses": {
          "200": { "description": "OK", "content": { "application/json": { "schema": { "$ref": "#/components/schemas/ResponsePendingTransmissionProcess" } } } }
        }
      }
    },
    "/configuration/parameters": {
      "get": {
        "tags": ["Configuration"],
        "summary": "Obtiene parámetros de configuración",
        "responses": {
          "200": { "description": "OK", "content": { "application/json": { "schema": { "$ref": "#/components/schemas/ResponseConfig" } } } }
        }
      }
    },
    "/configuration/promoter-duty": {
      "get": {
        "tags": ["Configuration"],
        "summary": "Obtiene deberes del promotor",
        "responses": {
          "200": { "description": "OK", "content": { "application/json": { "schema": { "$ref": "#/components/schemas/ResponsePromoterDutyList" } } } }
        }
      }
    },
    "/configuration/configuracion-inicial": {
      "get": {
        "tags": ["Configuration"],
        "summary": "Obtiene la configuración inicial del POS",
        "responses": {
          "200": { "description": "OK", "content": { "application/json": { "schema": { "$ref": "#/components/schemas/ResponseInitialConfiguration" } } } }
        }
      }
    },
    "/reports/day-closing-report/{fecha}": {
      "get": {
        "tags": ["Reports"],
        "summary": "Reporte de cierre diario",
        "parameters": [ { "name": "fecha", "in": "path", "required": true, "schema": { "type": "string" }, "description": "Fecha (YYYY-MM-DD)" } ],
        "responses": {
          "200": { "description": "OK", "content": { "application/json": { "schema": { "$ref": "#/components/schemas/ResponseDayClosingReport" } } } }
        }
      }
    },
    "/reports/fuel-report/{fecha}": {
      "get": {
        "tags": ["Reports"],
        "summary": "Reporte de combustible",
        "parameters": [ { "name": "fecha", "in": "path", "required": true, "schema": { "type": "string" }, "description": "Fecha (YYYY-MM-DD)" } ],
        "responses": {
          "200": { "description": "OK", "content": { "application/json": { "schema": { "$ref": "#/components/schemas/ResponseFuelReport" } } } }
        }
      }
    },
    "/sales/check-pending-sales": {
      "post": {
        "tags": ["Sales"],
        "summary": "Verifica ventas pendientes",
        "requestBody": { "required": true, "content": { "application/json": { "schema": { "$ref": "#/components/schemas/CheckPendingSalesRequest" } } } },
        "responses": {
          "200": { "description": "OK", "content": { "application/json": { "schema": { "$ref": "#/components/schemas/ResponsePendingSalesList" } } } }
        }
      }
    },
    "/sales/check-ready-sales": {
      "post": {
        "tags": ["Sales"],
        "summary": "Verifica ventas listas",
        "requestBody": { "required": true, "content": { "application/json": { "schema": { "$ref": "#/components/schemas/CheckReadySalesRequest" } } } },
        "responses": {
          "200": { "description": "OK", "content": { "application/json": { "schema": { "$ref": "#/components/schemas/ResponsePendingSalesList" } } } }
        }
      }
    },
    "/sales/datafono-cancellations-in-progress": {
      "post": {
        "tags": ["Sales"],
        "summary": "Cancelaciones en progreso de datáfono",
        "requestBody": { "required": true, "content": { "application/json": { "schema": { "$ref": "#/components/schemas/DatafonoCancellationsInProgressRequest" } } } },
        "responses": {
          "200": { "description": "OK", "content": { "application/json": { "schema": { "$ref": "#/components/schemas/ResponseDatafonoCancellationsInProgress" } } } }
        }
      }
    },
    "/sales/unresolved/attributes/{movementId}": {
      "get": {
        "tags": ["Sales"],
        "summary": "Atributos no resueltos de una venta",
        "parameters": [ { "name": "movementId", "in": "path", "required": true, "schema": { "type": "integer" }, "description": "ID del movimiento" } ],
        "responses": {
          "200": { "description": "OK", "content": { "application/json": { "schema": { "$ref": "#/components/schemas/ResponseUnresolvedSaleAttributes" } } } }
        }
      }
    },
    "/sales/movements/state/{movementId}": {
      "patch": {
        "tags": ["Sales"],
        "summary": "Actualiza estado de un movimiento",
        "parameters": [ { "name": "movementId", "in": "path", "required": true, "schema": { "type": "integer" }, "description": "ID del movimiento" } ],
        "requestBody": { "required": true, "content": { "application/json": { "schema": { "$ref": "#/components/schemas/UpdateMovementStateRequest" } } } },
        "responses": {
          "200": { "description": "OK", "content": { "application/json": { "schema": { "$ref": "#/components/schemas/ResponseUpdateMovementStateResult" } } } }
        }
      }
    },
    "/sales/assign-customer-data": {
      "post": {
        "tags": ["Sales"],
        "summary": "Asigna datos de cliente a la venta",
        "requestBody": { "required": true, "content": { "application/json": { "schema": { "$ref": "#/components/schemas/AssignCustomerDataRequest" } } } },
        "responses": {
          "200": { "description": "OK", "content": { "application/json": { "schema": { "$ref": "#/components/schemas/ResponseAssignCustomerDataResult" } } } }
        }
      }
    },
    "/sales/update-client-movement": {
      "patch": {
        "tags": ["Sales"],
        "summary": "Actualiza movimiento del cliente",
        "requestBody": { "required": true, "content": { "application/json": { "schema": { "$ref": "#/components/schemas/UpdateClientMovementRequest" } } } },
        "responses": {
          "200": { "description": "OK", "content": { "application/json": { "schema": { "$ref": "#/components/schemas/ResponseUpdateClientMovementResult" } } } }
        }
      }
    },
    "/sales/get-pending-sale-datafono/{id_transaccion}": {
      "get": {
        "tags": ["Sales"],
        "summary": "Obtiene estado de transacción de datáfono",
        "parameters": [ { "name": "id_transaccion", "in": "path", "required": true, "schema": { "type": "integer" }, "description": "ID de la transacción de datáfono" } ],
        "responses": {
          "200": { "description": "OK", "content": { "application/json": { "schema": { "$ref": "#/components/schemas/ResponsePendingSaleDatafono" } } } }
        }
      }
    },
    "/sales/update-payment-methods": {
      "patch": {
        "tags": ["Sales"],
        "summary": "Actualiza medios de pago de una venta",
        "requestBody": { "required": true, "content": { "application/json": { "schema": { "$ref": "#/components/schemas/UpdatePaymentMethodsRequest" } } } },
        "responses": {
          "200": { "description": "OK", "content": { "application/json": { "schema": { "$ref": "#/components/schemas/ResponseUpdatePaymentMethodsResult" } } } }
        }
      }
    },
    "/sales/fuel-entry-report": {
      "post": {
        "tags": ["Sales"],
        "summary": "Genera impresión de reporte de entrada de combustible",
        "requestBody": { "required": true, "content": { "application/json": { "schema": { "$ref": "#/components/schemas/FuelEntryReportRequest" } } } },
        "responses": {
          "200": { "description": "OK", "content": { "application/json": { "schema": { "$ref": "#/components/schemas/ResponseFuelEntryReportResult" } } } }
        }
      }
    },
    "/reports/closing-novelties": {
      "post": {
        "tags": ["Reports"],
        "summary": "Novedades de cierre diario",
        "requestBody": { "required": true, "content": { "application/json": { "schema": { "$ref": "#/components/schemas/DailyNoveltiesRequest" } } } },
        "responses": {
          "200": { "description": "OK", "content": { "application/json": { "schema": { "$ref": "#/components/schemas/ResponseDailyNovelties" } } } }
        }
      }
    }
  },
  "components": {
    "schemas": {
      "ResponseShift": {
        "type": "object",
        "properties": {
          "status": { "type": "integer" },
          "message": { "type": "string" },
          "process_date": { "type": "string" },
          "data": { "type": "object" }
        }
      },
      "ResponseShiftArray": {
        "allOf": [
          { "$ref": "#/components/schemas/ResponseShift" },
          { "type": "object", "properties": { "data": { "type": "array", "items": { "type": "object" } } } }
        ]
      },
      "PersonShift": {
        "type": "object",
        "properties": {
          "id": { "type": "integer" },
          "identificacion": { "type": "string" },
          "pin": { "type": "string" },
          "nombres": { "type": "string" },
          "apellidos": { "type": "string" },
          "perfiles_id": { "type": "integer" },
          "jornadas_id": { "type": "integer" }
        }
      },
      "PersonValidationRequest": {
        "type": "object",
        "properties": {
          "usuario": { "type": "string" },
          "clave": { "type": "string" },
          "tag": { "type": "string" }
        },
        "required": ["usuario", "clave"]
      },
      "ResponsePersonShift": {
        "allOf": [
          { "$ref": "#/components/schemas/ResponseShift" },
          { "type": "object", "properties": { "data": { "$ref": "#/components/schemas/PersonShift" } } }
        ]
      },

      "ResponseBase": {
        "type": "object",
        "properties": {
          "status": { "type": "integer" },
          "message": { "type": "string" },
          "process_date": { "type": "string" },
          "error": { "type": "string" }
        },
        "required": ["status", "process_date"]
      },

      "TransmissionProcessSummary": {
        "type": "object",
        "properties": {
          "processed": { "type": "integer" },
          "synchronized": { "type": "integer" },
          "failed": { "type": "integer" }
        }
      },
      "ResponsePendingTransmissionProcess": { "allOf": [ { "$ref": "#/components/schemas/ResponseBase" }, { "type": "object", "properties": { "data": { "$ref": "#/components/schemas/TransmissionProcessSummary" } } } ] },

      "PendingSale": {
        "type": "object",
        "properties": {
          "numero": { "type": "integer" },
          "razon_social": { "type": "string" },
          "nit": { "type": "string" },
          "cantidad": { "type": "number" },
          "precio": { "type": "number" },
          "total": { "type": "number" },
          "tipo": { "type": "string" },
          "recaudo": { "type": "integer" },
          "producto": { "type": "string" },
          "unidad_medida": { "type": "string" },
          "consecutivo": { "type": "integer" },
          "surtidor": { "type": "integer" },
          "cara": { "type": "integer" },
          "manguera": { "type": "integer" },
          "islas": { "type": "integer" },
          "atributos": { "type": "object" },
          "copia": { "type": "string" },
          "fecha": { "type": "string", "format": "date-time" },
          "id_promotor": { "type": "integer" },
          "operador": { "type": "string" },
          "identificacionPromotor": { "type": "string" },
          "identificacionProducto": { "type": "integer" },
          "medios_pagos": { "type": "object" },
          "id_tipo_venta": { "type": "integer" },
          "id_transmision": { "type": "integer" },
          "sincronizado": { "type": "integer" },
          "ind_pendiente_asignar_cliente": { "type": "boolean" },
          "proceso": { "type": "string" },
          "id_transaccion_datafono": { "type": "number" },
          "codigo_autorizacion_datafono": { "type": "string" },
          "id_transaccion_estado_datafono": { "type": "integer" },
          "descripcion_transaccion_estado_datafono": { "type": "string" },
          "ind_pendiente_resolver_datafono": { "type": "boolean" },
          "ind_pendiente_resolver_adblue": { "type": "boolean" },
          "estado_pagos": { "type": "string" },
          "integracion": { "type": "integer" }
        }
      },

      "CheckPendingSalesRequest": {
        "type": "object",
        "properties": {
          "journal_id": { "type": "integer" },
          "promoter_id": { "type": "integer" },
          "limit": { "type": "integer" }
        },
        "required": ["journal_id", "promoter_id", "limit"]
      },
      "CheckReadySalesRequest": {
        "type": "object",
        "properties": {
          "journal_id": { "type": "integer" },
          "promoter_id": { "type": "integer" },
          "limit": { "type": "integer" }
        },
        "required": ["journal_id", "promoter_id", "limit"]
      },
      "DatafonoCancellationsInProgressRequest": {
        "type": "object",
        "properties": {
          "id_movimiento": { "type": "integer" },
          "id_transaccion_operacion": { "type": "integer" },
          "id_transaccion_estado": { "type": "integer" }
        },
        "required": ["id_movimiento", "id_transaccion_operacion", "id_transaccion_estado"]
      },
      "UpdateMovementStateRequest": {
        "type": "object",
        "properties": { "estado_dian_id": { "type": "integer" } },
        "required": ["estado_dian_id"]
      },
      "AssignCustomerDataRequest": { "type": "object", "additionalProperties": true },
      "UpdateClientMovementRequest": {
        "type": "object",
        "properties": {
          "i_id_movimiento": { "type": "integer", "format": "int64" },
          "i_id_transmision": { "type": "integer", "format": "int64" },
          "i_sinconizacion": { "type": "integer" }
        },
        "required": ["i_id_movimiento", "i_id_transmision", "i_sinconizacion"]
      },

      "DatafonoCancellationsInProgress": { "type": "object", "properties": { "in_progress": { "type": "boolean" } } },
      "UnresolvedSaleAttributes": { "type": "object", "properties": { "atributos": { "type": "object" } } },
      "UpdateMovementStateResult": { "type": "object", "properties": { "updated": { "type": "boolean" } } },
      "AssignCustomerDataResult": { "type": "object", "properties": { "info": { "type": "object" } } },
      "UpdateClientMovementResult": { "type": "object", "properties": { "o_json_respuesta": { "type": "object" } } },

      "PendingSaleDatafono": {
        "type": "object",
        "properties": {
          "id_transaccion_estado": { "type": "integer" },
          "descripcion": { "type": "string" },
          "id_adquiriente": { "type": "integer" },
          "proveedor": { "type": "string" }
        }
      },
      "UpdatePaymentMethodsRequest": {
        "type": "object",
        "properties": {
          "identificadorMovimiento": { "type": "integer", "format": "int64" },
          "mediosDePagos": {
            "type": "array",
            "items": { "$ref": "#/components/schemas/PaymentMethodItem" }
          }
        }
      },
      "PaymentMethodItem": {
        "type": "object",
        "properties": {
          "ing_pago_datafono": { "type": "boolean" },
          "ct_medios_pagos_id": { "type": "integer", "format": "int64" },
          "valor_recibido": { "type": "number" },
          "valor_cambio": { "type": "number" },
          "valor_total": { "type": "number" },
          "numero_comprobante": { "type": "string" },
          "confirmacionBono": { "type": "boolean" }
        }
      },
      "UpdatePaymentMethodsResult": { "type": "object", "properties": { "info": { "type": "object" } } },

      "FuelEntryReportRequest": {
        "type": "object",
        "properties": {
          "numero_factura": { "type": "integer", "format": "int64" },
          "copia": { "type": "boolean" },
          "cola": { "type": "boolean" }
        },
        "required": ["numero_factura", "copia", "cola"]
      },
      "FuelEntryReportResult": { "type": "object", "properties": { "data": { "type": "object" } } },

      "ResponsePendingSalesList": { "allOf": [ { "$ref": "#/components/schemas/ResponseBase" }, { "type": "object", "properties": { "data": { "type": "array", "items": { "$ref": "#/components/schemas/PendingSale" } } } } ] },
      "ResponseDatafonoCancellationsInProgress": { "allOf": [ { "$ref": "#/components/schemas/ResponseBase" }, { "type": "object", "properties": { "data": { "$ref": "#/components/schemas/DatafonoCancellationsInProgress" } } } ] },
      "ResponseUnresolvedSaleAttributes": { "allOf": [ { "$ref": "#/components/schemas/ResponseBase" }, { "type": "object", "properties": { "data": { "$ref": "#/components/schemas/UnresolvedSaleAttributes" } } } ] },
      "ResponseUpdateMovementStateResult": { "allOf": [ { "$ref": "#/components/schemas/ResponseBase" }, { "type": "object", "properties": { "data": { "$ref": "#/components/schemas/UpdateMovementStateResult" } } } ] },
      "ResponseAssignCustomerDataResult": { "allOf": [ { "$ref": "#/components/schemas/ResponseBase" }, { "type": "object", "properties": { "data": { "$ref": "#/components/schemas/AssignCustomerDataResult" } } } ] },
      "ResponseUpdateClientMovementResult": { "allOf": [ { "$ref": "#/components/schemas/ResponseBase" }, { "type": "object", "properties": { "data": { "$ref": "#/components/schemas/UpdateClientMovementResult" } } } ] },
      "ResponsePendingSaleDatafono": { "allOf": [ { "$ref": "#/components/schemas/ResponseBase" }, { "type": "object", "properties": { "data": { "$ref": "#/components/schemas/PendingSaleDatafono" } } } ] },
      "ResponseUpdatePaymentMethodsResult": { "allOf": [ { "$ref": "#/components/schemas/ResponseBase" }, { "type": "object", "properties": { "data": { "$ref": "#/components/schemas/UpdatePaymentMethodsResult" } } } ] },
      "ResponseFuelEntryReportResult": { "allOf": [ { "$ref": "#/components/schemas/ResponseBase" }, { "type": "object", "properties": { "data": { "$ref": "#/components/schemas/FuelEntryReportResult" } } } ] },

      "EnvelopesTotalRequest": {
        "type": "object",
        "properties": { "journal_id": { "type": "integer" }, "promoter_id": { "type": "integer" } },
        "required": ["journal_id", "promoter_id"]
      },
      "TotalEnvelopes": { "type": "object", "properties": { "total": { "type": "number" } } },
      "ResponseTotalEnvelopes": { "allOf": [ { "$ref": "#/components/schemas/ResponseBase" }, { "type": "object", "properties": { "data": { "$ref": "#/components/schemas/TotalEnvelopes" } } } ] },

      "ErrorNotification": {
        "type": "object",
        "properties": {
          "detalle": { "type": "string" }
        }
      },
      "ResponseErrorNotificationList": {
        "allOf": [
          { "$ref": "#/components/schemas/ResponseBase" },
          { "type": "object", "properties": { "data": { "type": "array", "items": { "$ref": "#/components/schemas/ErrorNotification" } } } }
        ]
      },

      "Config": { "type": "object", "properties": { "tipo_autorizacion": { "type": "string" }, "solicitar_lecturas_tanques": { "type": "string" } } },
      "PromoterDuty": { "type": "object", "properties": { "personas_id": { "type": "integer" }, "nombre": { "type": "string" }, "estado": { "type": "string" }, "id_perfiles": { "type": "integer" }, "descripcion": { "type": "string" } } },
      "ResponseConfig": { "allOf": [ { "$ref": "#/components/schemas/ResponseBase" }, { "type": "object", "properties": { "data": { "$ref": "#/components/schemas/Config" } } } ] },
      "ResponsePromoterDutyList": { "allOf": [ { "$ref": "#/components/schemas/ResponseBase" }, { "type": "object", "properties": { "data": { "type": "array", "items": { "$ref": "#/components/schemas/PromoterDuty" } } } } ] },

      "InitialConfiguration": {
        "type": "object",
        "properties": {
          "equipo": { "type": "object", "additionalProperties": true },
          "jornada": { "type": "object", "additionalProperties": true },
          "empresas": { "type": "array", "items": { "type": "object", "additionalProperties": true } },
          "promotor": { "type": "object", "additionalProperties": true },
          "parametros": { "type": "object", "additionalProperties": true },
          "surtidores": { "type": "array", "items": { "type": "object", "additionalProperties": true } },
          "medios_pagos": { "type": "array", "items": { "type": "object", "additionalProperties": true } },
          "turno_activo": { "type": "boolean" },
          "surtidores_detalles": { "type": "array", "items": { "type": "object", "additionalProperties": true } }
        }
      },
      "ResponseInitialConfiguration": { "allOf": [ { "$ref": "#/components/schemas/ResponseBase" }, { "type": "object", "properties": { "data": { "$ref": "#/components/schemas/InitialConfiguration" } } } ] },

      "FuelReport": { "type": "object", "properties": { "total_ventas_combustible": { "type": "number" }, "total_ventas_canastilla": { "type": "number" }, "cantidad_ventas_combustible": { "type": "number" }, "cantidad_ventas_canastilla": { "type": "number" } } },
      "DayClosingReport": {
        "type": "object",
        "properties": {
          "totalVentasCombustible": { "type": "number" },
          "totalVentasCanastilla": { "type": "number" },
          "cantidadVentasCombustible": { "type": "number" },
          "cantidadVentasCanastilla": { "type": "number" },
          "cantidadVentasCDL": { "type": "number" },
          "totalVentasCDL": { "type": "number" },
          "ReporteMedios": { "type": "array", "items": { "$ref": "#/components/schemas/ReporteMedio" } }
        }
      },
      "ReporteMedio": { "type": "object", "properties": { "id": { "type": "integer" }, "descripcion": { "type": "string" }, "total": { "type": "number" }, "cantidad": { "type": "integer" } } },
      "ResponseFuelReport": { "allOf": [ { "$ref": "#/components/schemas/ResponseBase" }, { "type": "object", "properties": { "data": { "$ref": "#/components/schemas/FuelReport" } } } ] },
      "ResponseDayClosingReport": { "allOf": [ { "$ref": "#/components/schemas/ResponseBase" }, { "type": "object", "properties": { "data": { "$ref": "#/components/schemas/DayClosingReport" } } } ] },

      "DailyNoveltiesRequest": {
        "type": "object",
        "properties": {
          "ano": { "type": "integer" },
          "mes": { "type": "integer" },
          "dia": { "type": "integer" }
        },
        "required": ["ano", "mes", "dia"]
      },
      "DailyNovelties": {
        "type": "array",
        "items": { "type": "object", "additionalProperties": true }
      },
      "ResponseDailyNovelties": { "allOf": [ { "$ref": "#/components/schemas/ResponseBase" }, { "type": "object", "properties": { "data": { "$ref": "#/components/schemas/DailyNovelties" } } } ] },

      "OpeningShiftRequest": {
        "type": "object",
        "properties": {
          "usuario": { "type": "string" },
          "clave": { "type": "string" },
          "surtidores": { "type": "array", "items": { "type": "integer", "format": "int32" } },
          "fecha_inicio": { "type": "string" },
          "equipos_id": { "type": "integer", "format": "int32" },
          "empresas_id": { "type": "integer", "format": "int32" },
          "atributos": {
            "$ref": "#/components/schemas/OpeningAttributes"
          },
          "ajustePeriodico": { "type": "object" }
        },
        "required": ["usuario", "clave", "surtidores", "fecha_inicio", "equipos_id", "empresas_id", "atributos"]
      },
      "OpeningAttributes": {
        "type": "object",
        "properties": {
          "saldo": { "type": "integer" },
          "totalizadoresIniciales": {
            "type": "array",
            "items": { "$ref": "#/components/schemas/OpeningInitialTotals" }
          }
        }
      },
      "OpeningInitialTotals": {
        "type": "object",
        "properties": {
          "surtidor": { "type": "integer" },
          "cara": { "type": "integer" },
          "manguera": { "type": "integer" },
          "grado": { "type": "integer" },
          "isla": { "type": "integer" },
          "productoIdentificador": { "type": "integer" },
          "productoDescripcion": { "type": "string" },
          "familiaIdentificador": { "type": "integer" },
          "familiaDescripcion": { "type": "string" },
          "acumuladoVolumen": { "type": "integer" },
          "acumuladoVolumenReal": { "type": "integer" },
          "acumuladoVenta": { "type": "integer" },
          "factor_inventario": { "type": "integer" },
          "factor_volumen_parcial": { "type": "integer" },
          "factor_importe_parcial": { "type": "integer" },
          "factor_precio": { "type": "integer" },
          "precio": { "type": "number" }
        }
      }
    }
  }
}`
