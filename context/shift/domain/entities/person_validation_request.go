package entities

type PersonValidationRequest struct {
	Usuario string `json:"usuario" binding:"required"`
	Clave   string `json:"clave" binding:"required"`
	Tag     string `json:"tag"`
}
