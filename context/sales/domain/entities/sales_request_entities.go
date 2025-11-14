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
