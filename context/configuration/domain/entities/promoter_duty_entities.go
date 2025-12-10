package entities

type PromoterDuty struct {
	PersonasID             int    `json:"personas_id"`
	Nombre                 string `json:"nombre"`
	Estado                 string `json:"estado"`
	IdPerfiles             int    `json:"id_perfiles"`
	Descripcion            string `json:"descripcion"`
	IdentificacionPromotor string `json:"identificacion_promotor"`
	Jornada                int64  `json:"jornada"`
	FechaInicio            string `json:"fecha_inicio"`
}
