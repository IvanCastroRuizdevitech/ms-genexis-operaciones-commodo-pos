package constants

const (
    QUERY_GET_PERSON_SHIFT = "SELECT * FROM procesos.fnc_validacion_personal_eds($1, $2);"
    API_OPENING_SHIFT      = "http://localhost:10556/api/apertura/turno"
    QUERY_GET_DAILY_INCOME_MEASUREMENTS = "SELECT * FROM medidas_tanques_diarias mdt WHERE mdt.fecha_medida >= $1::timestamp AND mdt.fecha_medida <= $2::timestamp;"
)
