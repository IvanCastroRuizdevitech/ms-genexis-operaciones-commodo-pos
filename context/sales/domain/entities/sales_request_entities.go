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
