package constants

const (
	QUERY_GET_PERSON_SHIFT              = "SELECT * FROM procesos.fnc_validacion_personal_eds($1, $2, $3);"
	QUERY_VALIDATE_PERSON_SHIFT_ADMIN   = "SELECT * FROM procesos.fnc_validacion_personal_eds_admin($1, $2, $3, $4);"
	API_OPENING_SHIFT                   = "http://localhost:10556/api/apertura/turno"
	QUERY_GET_DAILY_INCOME_MEASUREMENTS = "SELECT * FROM medidas_tanques_diarias mdt WHERE mdt.fecha_medida >= $1::timestamp AND mdt.fecha_medida <= $2::timestamp;"
	QUERY_GET_FUEL_PUMPS                = "SELECT * FROM procesos.fnc_obtener_surtidores_por_politica();"
	QUERY_CREATE_ENVELOPE               = "select * from public.fnc_procesar_registro_sobre(i_sobres => $1::jsonb);"
	QUERY_GET_PERSON_BY_ID              = "select * from public.fnc_obtener_persona_por_id($1::bigint);"
)
