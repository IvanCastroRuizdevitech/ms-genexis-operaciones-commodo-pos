package entities

// TankBodega represents a tank/warehouse row from ct_bodegas.
type TankBodega struct {
	Id     float64 `json:"id"`
	Bodega string  `json:"bodega"`
}
