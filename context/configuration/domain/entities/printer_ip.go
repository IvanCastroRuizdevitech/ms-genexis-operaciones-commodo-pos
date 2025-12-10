package entities

type PrinterIPResult struct {
	Success bool             `json:"success,omitempty"`
	Data    PrinterIPPayload `json:"data,omitempty"`
}

type PrinterIPPayload struct {
	ID     int    `json:"id,omitempty"`
	Codigo string `json:"codigo,omitempty"`
	Valor  string `json:"valor,omitempty"`
}

// PrinterIPUpdateRequest is the payload to update the printer IP.
type PrinterIPUpdateRequest struct {
	IP string `json:"ip" binding:"required"` // allow IP or hostname; validation only checks presence
}

// PrinterIPUpdateResult carries the raw response from the update function and any parsed JSON if available.
type PrinterIPUpdateResult struct {
	Raw    string         `json:"raw_result"`
	Parsed map[string]any `json:"parsed,omitempty"`
}
