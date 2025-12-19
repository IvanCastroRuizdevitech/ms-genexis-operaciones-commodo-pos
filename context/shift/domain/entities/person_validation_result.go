package entities

type PersonValidationData struct {
	Id         int    `json:"id"`
	Nombre     string `json:"nombre"`
	PerfilesId int    `json:"perfiles_id"`
	TipoPerfil int    `json:"tipo_perfil"`
}

type PersonValidationResult struct {
	Success       bool                  `json:"success"`
	Authenticated bool                  `json:"authenticated"`
	Message       string                `json:"message"`
	Data          *PersonValidationData `json:"data,omitempty"`
}
