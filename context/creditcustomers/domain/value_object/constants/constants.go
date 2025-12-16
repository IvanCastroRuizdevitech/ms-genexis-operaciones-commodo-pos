package constants

const (
	QUERY_GET_DISPENSER_DETAILS             = "select * from public.fnc_obtener_surtidores_detalles();"
	QUERY_GET_DISPENSER_DETAILS_BY_FAMILIES = "select * from public.fnc_obtener_surtidores_detalles_por_familias($1);"
)
