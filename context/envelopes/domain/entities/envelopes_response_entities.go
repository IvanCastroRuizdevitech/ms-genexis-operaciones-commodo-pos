package entities

type TotalEnvelopes struct {
	Total float64 `json:"total"`
}

type EnvelopeCreate struct {
	Created      bool   `json:"created"`
	MessageError string `json:"message_error"`
}
