package constants

const (
    QUERY_CHECK_PENDING_SALES = "select * from fnc_consultar_ventas_pendientes(i_id_jornada => $1, i_id_responsable => $2, i_numero_registros => $3);"
    QUERY_CHECK_READY_SALES   = "select * from public.fnc_consultar_ventas(i_id_jornada => $1, i_id_responsable => $2, i_numero_registros => $3);"
    QUERY_CHECK_DATAFONO_CANCELLATIONS_IN_PROGRESS = "SELECT EXISTS (SELECT 1 FROM datafonos.transacciones AS t WHERE t.id_movimiento = $1 AND t.id_transaccion_operacion = $2 AND t.id_transaccion_estado = $3) AS in_progress;"
    QUERY_GET_UNRESOLVED_SALE_ATTRIBUTES = "select atributos from ct_movimientos cm where id = $1;"
    QUERY_UPDATE_MOVEMENT_STATE = "update facturacion_electronica.tbl_movimientos_facturas_electronicas SET estado_dian_id = $1, estado_descripcion = (select tede.descripcion from facturacion_electronica.tbl_estados_documentos_electronicos tede where tede.id = $2) where pos_movimiento_id = $3;"
    QUERY_ASSIGN_CUSTOMER_DATA = "select * from fnc_asignar_datos_cliente($1::json) as info;"
    QUERY_UPDATE_CLIENT_MOVEMENT = "select * from public.prc_registrar_cliente_movimiento($1,$2,$3,'{}'::json);"
    QUERY_GET_PENDING_SALE_DATAFONO = "SELECT td.id_transaccion_estado, td.descripcion, d.id_adquiriente, a.descripcion AS proveedor FROM datafonos.transacciones AS t INNER JOIN datafonos.transacciones_estado AS td ON t.id_transaccion_estado = td.id_transaccion_estado INNER JOIN datafonos.datafonos AS d ON t.id_datafono = d.id_datafono INNER JOIN datafonos.adquirientes AS a ON a.id_adquiriente = d.id_adquiriente WHERE t.id_transaccion = $1;"
)
