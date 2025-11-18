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
