package entities

type AssignTagTransmissionsRequest struct {
	Tag            string `json:"tag" binding:"required"`
	Identification string `json:"identification" binding:"required"`
	Medio          string `json:"medio" binding:"required"`
}

type AssignTagTransmissionsResult struct {
	TransmissionsRaw string `json:"transmissions_raw"`
}
