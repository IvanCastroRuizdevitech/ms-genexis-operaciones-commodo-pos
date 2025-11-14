package entities

// DayClosingReport represents a dynamic JSON object returned by the stored function
type DayClosingReport struct {
	TotalVentasCombustible    float64        `json:"totalVentasCombustible"`
	TotalVentasCanastilla     float64        `json:"totalVentasCanastilla"`
	CantidadVentasCombustible float64        `json:"cantidadVentasCombustible"`
	CantidadVentasCanastilla  float64        `json:"cantidadVentasCanastilla"`
	CantidadVentasCDL         float64        `json:"cantidadVentasCDL"`
	TotalVentasCDL            float64        `json:"totalVentasCDL"`
	ReporteMedios             []ReporteMedio `json:"ReporteMedios"`
}

type ReporteMedio struct {
	ID          int     `json:"id"`
	Descripcion string  `json:"descripcion"`
	Total       float64 `json:"total"`
	Cantidad    int     `json:"cantidad"`
}
