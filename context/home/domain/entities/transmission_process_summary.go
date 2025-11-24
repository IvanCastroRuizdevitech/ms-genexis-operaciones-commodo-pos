package entities

type TransmissionProcessSummary struct {
	Processed    int `json:"processed"`
	Synchronized int `json:"synchronized"`
	Failed       int `json:"failed"`
}
