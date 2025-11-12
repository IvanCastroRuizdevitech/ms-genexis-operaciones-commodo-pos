package entities

type PersonShift struct {
	Id             int    `json:"id"`
	Identificacion string `json:"identificacion"`
	Pin            string `json:"pin"`
	Nombres        string `json:"nombres"`
	Apellidos      string `json:"apellidos"`
	PerfilesId     int    `json:"perfiles_id"`
}

type ResponseShift struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	ProcessDate string      `json:"process_date"`
	Data        interface{} `json:"data"`
}
