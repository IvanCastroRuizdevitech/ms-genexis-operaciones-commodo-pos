package constants

const (
	QUERY_GET_DAY_CLOSING_REPORT = "select * from reporteria_cierres.fnc_consultar_cierre_dia(i_fecha => $1);"
	QUERY_GET_FUEL_REPORT        = `
SELECT
    COALESCE(SUM(((datos->>'lt_data')::json ->> 'TotalVentasSistema')::numeric), 0)        AS total_ventas_combustible,
    COALESCE(SUM(((datos->>'lt_data')::json ->> 'TotalVentasCanastilla')::numeric), 0)    AS total_ventas_canastilla,
    COALESCE(SUM(((datos->>'lt_data')::json ->> 'CantidadVentasSistema')::numeric), 0)    AS cantidad_ventas_combustible,
    COALESCE(SUM(((datos->>'lt_data')::json ->> 'NumeroVentasCanastilla')::numeric), 0)   AS cantidad_ventas_canastilla
FROM public.cierres c
WHERE datos IS NOT NULL
  AND (
        (
            (
                ((datos->>'lt_data')::json -> 'Turnos')::json -> 0
            ) ->> 'fecha'
        )::date = $1::date
      );
`
	QUERY_GET_DAILY_NOVELTIES     = "SELECT * FROM reporteria_cierres.obtener_novedades($1, $2, $3);"
	QUERY_CREATE_TANK_PRINT_EVENT = "CALL reporteria.insert_evento_impresion_lecturas_tanques($1::int[], NULL);"
	QUERY_GET_MOVEMENT_TYPES      = "select tm.id_tipo_movimiento as id, tm.descripcion from tipos_movimiento tm;"
	QUERY_GET_TANKS               = "select cb.id, cb.bodega from ct_bodegas cb;"
	QUERY_GET_SHIFT_SUMMARY       = "select * from reporteria_cierres.fnc_obtener_resumen_jornadas($1::int, $2::timestamp, $3::timestamp);"
	QUERY_GET_SHIFT_CONSOLIDATED  = "select * from reporteria_cierres.fnc_consulta_turnos_consolidado($1::timestamp, $2::timestamp);"
)
