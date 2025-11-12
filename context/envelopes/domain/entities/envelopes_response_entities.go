package entities

type TotalEnvelopes struct {
	Total float64 `json:"total"`
}

type ResponseEvelopesTotal struct {
	Status      int             `json:"status"`
	Message     string          `json:"message"`
	ProcessDate string          `json:"process_date"`
	Data        *TotalEnvelopes `json:"data"`
}
