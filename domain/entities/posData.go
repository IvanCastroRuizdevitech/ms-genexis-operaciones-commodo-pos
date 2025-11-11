package entities

type PosData struct {
	EquiposId         int64
	Host              string
	Isla              int64
	Atributos         map[string]interface{}
	CheckPosPrincipal bool
	Id_pos            string
}
