package entities

type PersonByIDRequest struct {
	PersonaID int64 `json:"persona_id" binding:"required"`
}

type PersonByIDData struct {
	ID         int64  `json:"id"`
	PerfilesID int64  `json:"perfiles_id"`
	Nombre     string `json:"nombre"`
	Usuario    string `json:"usuario"`
	Pin        string `json:"pin"`
}

type PersonByIDResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Total   *int            `json:"total,omitempty"`
	Data    *PersonByIDData `json:"data"`
	Error   *string         `json:"error,omitempty"`
}
