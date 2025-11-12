package constants

const (
	QUERY_GET_PERSON_SHIFT = "SELECT * FROM procesos.fnc_validacion_personal_eds($1, $2);"
	API_OPENING_SHIFT      = "http://localhost:10556/api/apertura/turno"
)
