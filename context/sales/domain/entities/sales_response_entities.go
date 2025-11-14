package entities

import "time"

type PendingSale struct {
    Numero                               int         `json:"numero" db:"numero"`
    RazonSocial                          string      `json:"razon_social" db:"razon_social"`
    Nit                                  string      `json:"nit" db:"nit"`
    Cantidad                             float64     `json:"cantidad" db:"cantidad"`
    Precio                               float64     `json:"precio" db:"precio"`
    Total                                float64     `json:"total" db:"total"`
    Tipo                                 string      `json:"tipo" db:"tipo"`
    Recaudo                              int         `json:"recaudo" db:"recaudo"`
    Producto                             string      `json:"producto" db:"producto"`
    UnidadMedida                         string      `json:"unidad_medida" db:"unidad_medida"`
    Consecutivo                          int64       `json:"consecutivo" db:"consecutivo"`
    Surtidor                             int         `json:"surtidor" db:"surtidor"`
    Cara                                 int         `json:"cara" db:"cara"`
    Manguera                             int         `json:"manguera" db:"manguera"`
    Islas                                int         `json:"islas" db:"islas"`
    Atributos                            any         `json:"atributos" db:"atributos"`
    Copia                                string      `json:"copia" db:"copia"`
    Fecha                                time.Time   `json:"fecha" db:"fecha"`
    IdPromotor                           int64       `json:"id_promotor" db:"id_promotor"`
    Operador                             string      `json:"operador" db:"operador"`
    IdentificacionPromotor               string      `json:"identificacionPromotor" db:"identificacionPromotor"`
    IdentificacionProducto               int64       `json:"identificacionProducto" db:"identificacionProducto"`
    MediosPagos                          any         `json:"medios_pagos" db:"medios_pagos"`
    IdTipoVenta                          int         `json:"id_tipo_venta" db:"id_tipo_venta"`
    IdTransmision                        int64       `json:"id_transmision" db:"id_transmision"`
    Sincronizado                         int         `json:"sincronizado" db:"sincronizado"`
    IndPendienteAsignarCliente           bool        `json:"ind_pendiente_asignar_cliente" db:"ind_pendiente_asignar_cliente"`
    Proceso                              string      `json:"proceso" db:"proceso"`
    IdTransaccionDatafono                float64     `json:"id_transaccion_datafono" db:"id_transaccion_datafono"`
    CodigoAutorizacionDatafono           string      `json:"codigo_autorizacion_datafono" db:"codigo_autorizacion_datafono"`
    IdTransaccionEstadoDatafono          int16       `json:"id_transaccion_estado_datafono" db:"id_transaccion_estado_datafono"`
    DescripcionTransaccionEstadoDatafono string      `json:"descripcion_transaccion_estado_datafono" db:"descripcion_transaccion_estado_datafono"`
    IndPendienteResolverDatafono         bool        `json:"ind_pendiente_resolver_datafono" db:"ind_pendiente_resolver_datafono"`
    IndPendienteResolverAdblue           bool        `json:"ind_pendiente_resolver_adblue" db:"ind_pendiente_resolver_adblue"`
    EstadoPagos                          string      `json:"estado_pagos" db:"estado_pagos"`
    Integracion                          int64       `json:"integracion" db:"integracion"`
}

type DatafonoCancellationsInProgress struct {
    InProgress bool `json:"in_progress" db:"in_progress"`
}
