package constants

const (
	QUERY_UPDATE_COMANDA_STATUS = "select * from restaurante.fnc_actualizar_estado_comanda($1::bigint, $2::smallint);"
)
