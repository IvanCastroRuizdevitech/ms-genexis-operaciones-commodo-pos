package entities

// TankPrintEventRequest represents the payload to trigger a tank inventory print event.
type TankPrintEventRequest struct {
	TankIDs []int `json:"tank_ids"`
}

// TankPrintEventResult represents the outcome of triggering the print event.
type TankPrintEventResult struct {
	Created bool `json:"created"`
}
