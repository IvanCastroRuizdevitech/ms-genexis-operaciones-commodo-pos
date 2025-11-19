package entities

type CheckPendingSalesRequest struct {
    JournalId   int `json:"journal_id" binding:"required"`
    PromoterId  int `json:"promoter_id" binding:"required"`
    Limit       int `json:"limit" binding:"required"`
}

type CheckReadySalesRequest struct {
    JournalId   int `json:"journal_id" binding:"required"`
    PromoterId  int `json:"promoter_id" binding:"required"`
    Limit       int `json:"limit" binding:"required"`
}

type DatafonoCancellationsInProgressRequest struct {
    MovementId             int `json:"id_movimiento" binding:"required"`
    TransactionOperationId int `json:"id_transaccion_operacion" binding:"required"`
    TransactionStatusId    int `json:"id_transaccion_estado" binding:"required"`
}

type UpdateMovementStateRequest struct {
    EstadoDianId int `json:"estado_dian_id" binding:"required"`
}

// AssignCustomerDataRequest represents a passthrough JSON payload
// that will be forwarded to the database function as a single JSON argument.
// Using a map preserves flexibility for varying parameter sets.
type AssignCustomerDataRequest map[string]any

// UpdateClientMovementRequest payload for updating client movement through stored procedure
// public.prc_registrar_cliente_movimiento(?,?,?,'{}'::json)
// (IN i_id_movimiento bigint, IN i_id_transmision bigint, IN i_sinconizacion integer)
type UpdateClientMovementRequest struct {
    MovementId     int64 `json:"i_id_movimiento" binding:"required"`
    TransmissionId int64 `json:"i_id_transmision" binding:"required"`
    Synchronization int  `json:"i_sinconizacion" binding:"required"`
}

type UpdatePaymentMethodsRequest struct {
    IdMovimiento int64                 `json:"identificadorMovimiento,omitempty"`
    MediosPago   []PaymentMethodItem   `json:"mediosDePagos,omitempty"`
}

type PaymentMethodItem struct {
    IngPagoDatafono   bool    `json:"ing_pago_datafono"`
    IdMedio           int64   `json:"ct_medios_pagos_id"`
    ValorRecibido     float64 `json:"valor_recibido"`
    ValorCambio       float64 `json:"valor_cambio"`
    ValorTotal        float64 `json:"valor_total"`
    NumeroComprobante string  `json:"numero_comprobante"`
    ConfirmacionBono  bool    `json:"confirmacionBono"`
}

// FuelEntryReportRequest payload
// Maps to procesos.fnc_re_imprimir_factura_entrada(numero_factura, copia, cola)
type FuelEntryReportRequest struct {
    NumeroFactura int64 `json:"numero_factura" binding:"required"`
    Copia         bool  `json:"copia" binding:"required"`
    Cola          bool  `json:"cola" binding:"required"`
}
