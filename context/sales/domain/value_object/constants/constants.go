package constants

const (
    QUERY_CHECK_PENDING_SALES = "select * from fnc_consultar_ventas_pendientes(i_id_jornada => $1, i_id_responsable => $2, i_numero_registros => $3);"
)

