package entities

// FuelReport represents the aggregated sales summary obtained from cierres.
type FuelReport struct {
	TotalVentasCombustible    float64 `json:"total_ventas_combustible"`
	TotalVentasCanastilla     float64 `json:"total_ventas_canastilla"`
	CantidadVentasCombustible float64 `json:"cantidad_ventas_combustible"`
	CantidadVentasCanastilla  float64 `json:"cantidad_ventas_canastilla"`
}
