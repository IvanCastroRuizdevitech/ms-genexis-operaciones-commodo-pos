package entities

type DataResponse struct {
}

type ResponseShift struct {
	Status      int           `json:"status"`
	Message     string        `json:"mensaje"`
	ProcessDate string        `json:"fechaProceso"`
	Data        *DataResponse `json:"data"`
}
