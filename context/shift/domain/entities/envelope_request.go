package entities

type EnvelopeRequest struct {
	IdentifierCompany      int     `json:"identificadorEmpresa" binding:"required"`
	Date                   string  `json:"fecha" binding:"required"`
	IdentifierPromoter     int     `json:"identificadorPromotor" binding:"required"`
	IdentifierDevice       int     `json:"identificadorEquipo" binding:"required"`
	IdentifierGroupJournal int     `json:"identificadorGrupoJornada" binding:"required"`
	Total                  float64 `json:"total" binding:"required"`
	RemoteID               int     `json:"remoto_id" binding:"required"`
	Attributes             any     `json:"atributos" binding:"required"`
}

type EnvelopeCreate struct {
	Created      bool   `json:"created"`
	MessageError string `json:"message_error"`
}
