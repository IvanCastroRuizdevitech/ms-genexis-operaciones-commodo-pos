package entities

import (
	"encoding/json"
	"time"
)

// DispenserDetail represents the active dispenser detail row returned by
// genexis_operaciones.fnc_obtener_surtidores_detalles().
type DispenserDetail struct {
	SurtidoresDetallesID int             `json:"surtidores_detalles_id" db:"surtidores_detalles_id"`
	SurtidoresID         int             `json:"surtidores_id" db:"surtidores_id"`
	ProductosID          int             `json:"productos_id" db:"productos_id"`
	Estado               json.RawMessage `json:"estado" db:"estado"`
	DescripcionProducto  string          `json:"descripcion_producto" db:"descripcion_producto"`
	PrecioProducto       float64         `json:"precio_producto" db:"precio_producto"`
	FamiliaCodigo        string          `json:"familia_codigo" db:"familia_codigo"`
	FamiliaID            int             `json:"familia_id" db:"familia_id"`
}

type PendingSale struct {
	Numero                               int       `json:"numero" db:"numero"`
	RazonSocial                          string    `json:"razon_social" db:"razon_social"`
	Nit                                  string    `json:"nit" db:"nit"`
	Cantidad                             float64   `json:"cantidad" db:"cantidad"`
	Precio                               float64   `json:"precio" db:"precio"`
	Total                                float64   `json:"total" db:"total"`
	Tipo                                 string    `json:"tipo" db:"tipo"`
	Recaudo                              int       `json:"recaudo" db:"recaudo"`
	Producto                             string    `json:"producto" db:"producto"`
	UnidadMedida                         string    `json:"unidad_medida" db:"unidad_medida"`
	Consecutivo                          int64     `json:"consecutivo" db:"consecutivo"`
	Surtidor                             int       `json:"surtidor" db:"surtidor"`
	Cara                                 int       `json:"cara" db:"cara"`
	Manguera                             int       `json:"manguera" db:"manguera"`
	Islas                                int       `json:"islas" db:"islas"`
	Atributos                            any       `json:"atributos" db:"atributos"`
	Copia                                string    `json:"copia" db:"copia"`
	Fecha                                time.Time `json:"fecha" db:"fecha"`
	IdPromotor                           int64     `json:"id_promotor" db:"id_promotor"`
	Operador                             string    `json:"operador" db:"operador"`
	IdentificacionPromotor               string    `json:"identificacionPromotor" db:"identificacionPromotor"`
	IdentificacionProducto               int64     `json:"identificacionProducto" db:"identificacionProducto"`
	MediosPagos                          any       `json:"medios_pagos" db:"medios_pagos"`
	IdTipoVenta                          int       `json:"id_tipo_venta" db:"id_tipo_venta"`
	IdTransmision                        int64     `json:"id_transmision" db:"id_transmision"`
	Sincronizado                         int       `json:"sincronizado" db:"sincronizado"`
	IndPendienteAsignarCliente           bool      `json:"ind_pendiente_asignar_cliente" db:"ind_pendiente_asignar_cliente"`
	Proceso                              string    `json:"proceso" db:"proceso"`
	IdTransaccionDatafono                float64   `json:"id_transaccion_datafono" db:"id_transaccion_datafono"`
	CodigoAutorizacionDatafono           string    `json:"codigo_autorizacion_datafono" db:"codigo_autorizacion_datafono"`
	IdTransaccionEstadoDatafono          int16     `json:"id_transaccion_estado_datafono" db:"id_transaccion_estado_datafono"`
	DescripcionTransaccionEstadoDatafono string    `json:"descripcion_transaccion_estado_datafono" db:"descripcion_transaccion_estado_datafono"`
	IndPendienteResolverDatafono         bool      `json:"ind_pendiente_resolver_datafono" db:"ind_pendiente_resolver_datafono"`
	IndPendienteResolverAdblue           bool      `json:"ind_pendiente_resolver_adblue" db:"ind_pendiente_resolver_adblue"`
	EstadoPagos                          string    `json:"estado_pagos" db:"estado_pagos"`
	Integracion                          int64     `json:"integracion" db:"integracion"`
}

type DatafonoCancellationsInProgress struct {
	InProgress bool `json:"in_progress" db:"in_progress"`
}

type UnresolvedSaleAttributes struct {
	Atributos any `json:"atributos" db:"atributos"`
}

type UpdateMovementStateResult struct {
	Updated bool `json:"updated" db:"updated"`
}

// AssignCustomerDataResult maps the result from
// fnc_asignar_datos_cliente(... ) AS info
type AssignCustomerDataResult struct {
	Info any `json:"info" db:"info"`
}

// UpdateClientMovementResult maps the OUT/INOUT json response from the stored routine
// prc_registrar_cliente_movimiento(...)
// The column name is assumed as "o_json_respuesta" when selecting the routine result.
type UpdateClientMovementResult struct {
	Json any `json:"o_json_respuesta" db:"o_json_respuesta"`
}

// PendingSaleDatafono maps the pending sale status for a datafono transaction
// returned by the query over datafonos.transacciones and related tables.
type PendingSaleDatafono struct {
	IdTransaccionEstado int    `json:"id_transaccion_estado" db:"id_transaccion_estado"`
	Descripcion         string `json:"descripcion" db:"descripcion"`
	IdAdquiriente       int    `json:"id_adquiriente" db:"id_adquiriente"`
	Proveedor           string `json:"proveedor" db:"proveedor"`
}

// UpdatePaymentMethodsResult maps the boolean result aliased as "info"
// from fnc_actualizar_medios_de_pagos($1::json) and includes
// the movement identifier and the payment methods payload shape.
type UpdatePaymentMethodsResult struct {
	// Database function result (true/false)
	Info any `json:"info" db:"info"`
}

type ReprintSaleResult struct {
	Info any `json:"info" db:"info"`
}

// Respuesta completa de la función PL/pgSQL
type FuelEntryReportResult struct {
	Codigo  int           `json:"codigo"`
	Mensaje string        `json:"mensaje"`
	Result  FacturaResult `json:"result"`
}

// "result"
type FacturaResult struct {
	Data          FacturaData `json:"data"`
	FechaRegistro string      `json:"fecha_registro"`
	JornadaID     int         `json:"jornada_id"` // si en el JSON viene como texto cambia a string
	PosID         int         `json:"pos_id"`
	Promotor      string      `json:"promotor"`
	Copia         any         `json:"copia"` // no sabemos el tipo exacto -> interface{}
}

// "data"
type FacturaData struct {
	InformacionGeneral      InformacionGeneral     `json:"INFORMACION_GENERAL"`
	MedidasIniciales        []MedidaInicial        `json:"MEDIDAS_INICIALES"`
	ProductosSeleccionados  []ProductoSeleccionado `json:"PRODUCTOS_SELECIONADOS"`
	SolicitudMedidasFinales []MedidaFinal          `json:"SOLICITUD_MEDIDAS_FINALES"`
	TanquesSeleccionados    []TanqueSeleccionado   `json:"TANQUES_SELECCIONADOS"`
	FechaFin                string                 `json:"fechaFin"`
	FechaInicio             string                 `json:"fechaInicio"`
	FechaTransaccion        string                 `json:"fechaTransaccion"`
	LecturaVeeder           any                    `json:"lecturaVeeder"` // viene null
	LecturaVeederFinal      LecturaVeederFinal     `json:"lecturaVeederFinal"`
	DiferenciaGalones       float64                `json:"diferenciaGalones"`
	DiferenciaVentas        float64                `json:"diferenciaVentas"`
}

// "INFORMACION_GENERAL"
type InformacionGeneral struct {
	Documento string `json:"DOCUMENTO"`
	Placa     string `json:"PLACA"`
}

// Elementos de "MEDIDAS_INICIALES"
type MedidaInicial struct {
	Agua                 float64 `json:"agua"`
	Altura               float64 `json:"altura"`
	Galones              float64 `json:"galones"`
	IdentificacionTanque int     `json:"identificacionTanque"`
	IdentificadorTanque  string  `json:"identificadorTanque"`
}

// Elementos de "PRODUCTOS_SELECIONADOS"
type ProductoSeleccionado struct {
	IdentificadorProducto int `json:"identificadorProducto"`
	IdentificacionTanque  int `json:"identificacionTanque"`
	IdentificadorTanque   int `json:"identificadorTanque"`
}

// Elementos de "SOLICITUD_MEDIDAS_FINALES"
type MedidaFinal struct {
	Agua                 float64 `json:"agua"`
	Altura               float64 `json:"altura"`
	Galones              float64 `json:"galones"`
	IdentificacionTanque int     `json:"identificacionTanque"`
	IdentificadorTanque  string  `json:"identificadorTanque"`
}

// Elementos de "TANQUES_SELECCIONADOS"
type TanqueSeleccionado struct {
	Cantidad              string `json:"cantidad"` // viene como texto en el JSON (->>)
	IdentificacionTanque  int    `json:"identificacionTanque"`
	IdentificadorProducto int    `json:"identificadorProducto"`
	IdentificadorTanque   int    `json:"identificadorTanque"`
	ProductoDesc          string `json:"productoDesc"`
	ProductoPrecio        int    `json:"productoPrecio"`
}

// Objeto "lecturaVeederFinal"
type LecturaVeederFinal struct {
	Agua    float64 `json:"agua"`
	Altura  float64 `json:"altura"`
	Volumen float64 `json:"volumen"`
}
