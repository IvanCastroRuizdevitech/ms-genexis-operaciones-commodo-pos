package entities

type PersonValidationData struct {
	Id           int64  `json:"id"`
	Nombre       string `json:"nombre"`
	PerfilesId   int64  `json:"perfiles_id"`
	TipoPerfil   int    `json:"tipo_perfil"`
	JornadasId   int64  `json:"jornadas_id"`
}

type PersonValidationResult struct {
	Success       bool                  `json:"success"`
	Authenticated bool                  `json:"authenticated"`
	Message       string                `json:"message"`
	Data          *PersonValidationData `json:"data,omitempty"`
}
