package constants

const (
	QUERY_GET_ALL_CONFIGURATION         = "select * from genexis_operaciones.fnc_obtener_configuracion();"
	QUERY_GET_PROMOTER_DUTY             = "select * from genexis_operaciones.fnc_get_promoter_duty();"
	QUERY_GET_INITIAL_CONFIGURATION_POS = "select * from genexis_operaciones.fnc_obtener_configuracion_pos();"
	QUERY_GET_PRINTER_IP                = "select * from public.fnc_obtener_ip_impresora();"
	QUERY_GET_UNIFIED_CONSECUTIVES      = "select * from public.fnc_obtener_consecutivos_unificados();"
	QUERY_UPDATE_PRINTER_IP             = "select * from public.fnc_actualizar_ip_impresora($1);"
	QUERY_VALIDATE_ADMIN_PERSON         = "select * from procesos.fnc_validacion_personal_eds_admin($1, $2, $3);"
	QUERY_PROCESS_NOTIFICATION          = "call public.prc_procesar_notificacion($1::bigint, $2::text, $3::bool, $4::json);"
	MICRO_PRINTER_URL                   = "http://127.0.0.1:8054/printer"
)
