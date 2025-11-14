package constants

const (
    QUERY_GET_DAY_CLOSING_REPORT = "select * from reporteria_cierres.fnc_consultar_cierre_dia(i_fecha => $1);"
)

