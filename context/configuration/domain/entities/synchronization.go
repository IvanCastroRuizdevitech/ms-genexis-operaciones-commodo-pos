package entities

import "time"

type SynchronizationQuery struct {
	IDSincronizacion int       `query:"id_sincronizacion" validate:"required,int"`
	FechaInicio      time.Time `query:"fecha_inicio" validate:"required,datetime"`
	FechaFin         time.Time `query:"fecha_fin" validate:"required,datetime"`
}

type SynchronizationDetail struct {
	ID               int    `json:"id"`
	IDNotificaciones int    `json:"id_notificaciones"`
	Descripcion      string `json:"descripcion"`
	Logger           string `json:"logger"`
	Fecha            string `json:"fecha"`
}
