package constants

const (
	QUERY_GET_DISPENSER_DETAILS               = "select * from public.fnc_obtener_surtidores_detalles();"
	QUERY_GET_DISPENSER_DETAILS_BY_FAMILIES   = "select * from public.fnc_obtener_surtidores_detalles_por_familias($1);"
	QUERY_GET_IDENTIFIER_TYPES               = "select * from clientes_credito.tbl_tipos_identificadores tti;"
	QUERY_GET_PRICE_FAMILIES                  = "select * from fnc_obtener_familia_precios() familiasPrecios;"
	QUERY_INSERT_PRE_AUTHORIZATION_CUSTOMER   = "select * from public.fnc_insertar_pre_autorizacion_cliente($1::jsonb) respuesta;"
	QUERY_UPDATE_TRANSACTION_BY_AUTHORIZATION = "" +
		"select * from public.fnc_actualizar_transaccion_por_autorizacion(" +
		"$1::uuid, " +
		"$2::integer, " +
		"$3::integer, " +
		"$4::integer, " +
		"$5::text, " +
		"$6::text, " +
		"$7::numeric, " +
		"$8::numeric, " +
		"$9::text, " +
		"$10::text, " +
		"$11::json, " +
		"$12::smallint, " +
		"$13::text, " +
		"$14::text, " +
		"$15::integer" +
		");"
)
