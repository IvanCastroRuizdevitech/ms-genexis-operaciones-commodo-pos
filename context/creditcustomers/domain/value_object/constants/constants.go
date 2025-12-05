package constants

const (
	QueryGetIdentifierTypes     = "select * from clientes_credito.fnc_obtener_tipos_identificadores();"
	QueryGetPriceFamilies       = "select * from fnc_obtener_familia_precios() familiasPrecios;"
	QueryInsertPreAuthorization = "select * from public.fnc_insertar_pre_autorizacion_cliente($1::jsonb) respuesta;"
)
