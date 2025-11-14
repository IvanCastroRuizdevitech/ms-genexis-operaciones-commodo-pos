package constants

const (
    QUERY_CHECK_PENDING_SALES = "select * from fnc_consultar_ventas_pendientes(i_id_jornada => $1, i_id_responsable => $2, i_numero_registros => $3);"
    QUERY_CHECK_READY_SALES   = "select * from public.fnc_consultar_ventas(i_id_jornada => $1, i_id_responsable => $2, i_numero_registros => $3);"
    QUERY_CHECK_DATAFONO_CANCELLATIONS_IN_PROGRESS = "SELECT EXISTS (SELECT 1 FROM datafonos.transacciones AS t WHERE t.id_movimiento = $1 AND t.id_transaccion_operacion = $2 AND t.id_transaccion_estado = $3) AS in_progress;"
    QUERY_GET_UNRESOLVED_SALE_ATTRIBUTES = "select atributos from ct_movimientos cm where id = $1;"
)
