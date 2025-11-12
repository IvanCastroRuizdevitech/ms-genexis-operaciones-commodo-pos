package entities

type EnvelopesTotalRequest struct {
	JournalId  int `json:"journal_id" binding:"required"`
	PromoterId int `json:"promoter_id" binding:"required"`
}

type EnvelopeRequest struct {
	IdentifierCompany      int     `json:"identificadorEmpresa" binding:"required"`
	Date                   string  `json:"fecha" binding:"required"`
	IdentifierPromoter     int     `json:"identificadorPromotor" binding:"required"`
	Total                  float64 `json:"total" binding:"required"`
	IdentifierDevice       int     `json:"identificadorEquipo" binding:"required"`
	RemoteId               int     `json:"remoto_id" binding:"required"`
	AttributeResponsible   any     `json:"atributos" binding:"required"`
	IdentifierGroupJournal int     `json:"identificadorGrupoJornada" binding:"required"`
}
