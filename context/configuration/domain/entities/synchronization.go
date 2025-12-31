package entities

type SynchronizationQuery struct {
	IDSincronizacion int    `query:"id_sincronizacion" validate:"required,int"`
	FechaInicio      string `query:"fecha_inicio" validate:"required,string"`
	FechaFin         string `query:"fecha_fin" validate:"required,string"`
}

type SynchronizationDetail struct {
	ID               int    `json:"id"`
	IDNotificaciones int    `json:"id_notificaciones"`
	Descripcion      string `json:"descripcion"`
	Logger           string `json:"logger"`
	Fecha            string `json:"fecha"`
}
