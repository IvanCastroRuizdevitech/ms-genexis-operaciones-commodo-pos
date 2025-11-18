package entities

// DailyNovelties represents the dynamic result set returned by the obtener_novedades function.
type DailyNovelties []map[string]interface{}

// DailyNoveltiesRequest holds the input payload for the daily novelties endpoint.
type DailyNoveltiesRequest struct {
	Ano int `json:"ano"`
	Mes int `json:"mes"`
	Dia int `json:"dia"`
}
