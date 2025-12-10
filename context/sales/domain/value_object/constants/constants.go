package constants

const (
	QUERY_CHECK_PENDING_SALES                      = "select * from public.fnc_consultar_ventas_pendientes(i_id_jornada => $1, i_id_responsable => $2, i_numero_registros => $3);"
	QUERY_CHECK_READY_SALES                        = "select * from public.fnc_consultar_ventas(i_id_jornada => $1, i_id_responsable => $2, i_numero_registros => $3);"
	QUERY_CHECK_DATAFONO_CANCELLATIONS_IN_PROGRESS = "SELECT EXISTS (SELECT 1 FROM datafonos.transacciones AS t WHERE t.id_movimiento = $1 AND t.id_transaccion_operacion = $2 AND t.id_transaccion_estado = $3) AS in_progress;"
	QUERY_GET_UNRESOLVED_SALE_ATTRIBUTES           = "select atributos from ct_movimientos cm where id = $1;"
	QUERY_UPDATE_MOVEMENT_STATE                    = "update facturacion_electronica.tbl_movimientos_facturas_electronicas SET estado_dian_id = $1, estado_descripcion = (select tede.descripcion from facturacion_electronica.tbl_estados_documentos_electronicos tede where tede.id = $2) where pos_movimiento_id = $3;"
	QUERY_ASSIGN_CUSTOMER_DATA                     = "select * from public.fnc_asignar_datos_cliente($1::json) as info;"
	QUERY_UPDATE_CLIENT_MOVEMENT                   = "select * from public.prc_registrar_cliente_movimiento($1,$2,$3,'{}'::json);"
	QUERY_GET_PENDING_SALE_DATAFONO                = "SELECT td.id_transaccion_estado, td.descripcion, d.id_adquiriente, a.descripcion AS proveedor FROM datafonos.transacciones AS t INNER JOIN datafonos.transacciones_estado AS td ON t.id_transaccion_estado = td.id_transaccion_estado INNER JOIN datafonos.datafonos AS d ON t.id_datafono = d.id_datafono INNER JOIN datafonos.adquirientes AS a ON a.id_adquiriente = d.id_adquiriente WHERE t.id_transaccion = $1;"
	QUERY_UPDATE_PAYMENT_METHODS                   = "select * from public.fnc_actualizar_medios_de_pagos($1::json) as info;"
	QUERY_REPRINT_SALE                             = "select * from public.reimpresion($1) as info;"
	QUERY_FUEL_ENTRY_REPORT                        = "select * from procesos.fnc_re_imprimir_factura_entrada($1::bigint, $2::boolean, $3::boolean) as data;"
	QUERY_GET_DISPENSER_DETAILS                    = "select * from genexis_operaciones.fnc_obtener_surtidores_detalles();"
	QUERY_UPSERT_ACTIVE_DISPENSER_FACE_TRANSACTION = "select * from genexis_operaciones.fnc_transaccion_surtidor_cara_vigente($1::int, $2::int, $3::varchar, $4::int, $5::int, $6::numeric, $7::numeric, $8::json, $9::int);"
)
