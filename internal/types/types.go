package types

// PrintRequest defines the expected POST body for /printer.
type PrintRequest struct {
	Template  []string       `json:"template"`
	Port      int            `json:"port"`
	Host      string         `json:"host"`
	ExtraData map[string]any `json:"extraData,omitempty"`
}

// PrintResponse is the unified HTTP response payload.
type PrintResponse struct {
	State   string `json:"state"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}
