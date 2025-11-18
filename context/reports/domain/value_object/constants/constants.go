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
	QUERY_GET_DAILY_NOVELTIES = "SELECT * FROM reporteria_cierres.obtener_novedades($1, $2, $3);"
)
