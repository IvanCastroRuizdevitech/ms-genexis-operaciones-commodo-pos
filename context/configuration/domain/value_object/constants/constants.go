package constants

const (
	QUERY_GET_ALL_CONFIGURATION         = "select * from genexis_operaciones.fnc_obtener_configuracion();"
	QUERY_GET_PROMOTER_DUTY             = "select * from genexis_operaciones.fnc_get_promoter_duty();"
	QUERY_GET_INITIAL_CONFIGURATION_POS = "select * from genexis_operaciones.fnc_obtener_configuracion_pos();"
	QUERY_GET_PRINTER_IP                = "select * from public.fnc_obtener_ip_impresora();"
	QUERY_UPDATE_PRINTER_IP             = "select * from public.fnc_actualizar_ip_impresora($1);"
	MICRO_PRINTER_URL                   = "http://127.0.0.1:8054/printer"
)
