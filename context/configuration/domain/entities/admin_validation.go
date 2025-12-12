package entities

// AdminValidationRequest carries the credentials or tag used to validate an admin user.
type AdminValidationRequest struct {
	Usuario string `json:"usuario" binding:"required_without=Tag"`
	Clave   string `json:"clave" binding:"required_without=Tag"`
	Tag     string `json:"tag" binding:"required_without_all=Usuario Clave"`
}

// AdminValidationResult is the JSON structure returned by the database function.
type AdminValidationResult struct {
	Success       bool   `json:"success"`
	Authenticated bool   `json:"authenticated"`
	Message       string `json:"message"`
}
