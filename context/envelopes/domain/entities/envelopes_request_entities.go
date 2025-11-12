package entities

type EnvelopesTotalRequest struct {
	JournalId  int `json:"journal_id" binding:"required"`
	PromoterId int `json:"promoter_id" binding:"required"`
}
