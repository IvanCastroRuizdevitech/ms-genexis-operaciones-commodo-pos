package constants

const QUERY_LOAD_ERROR_NOTIFICATION = `
SELECT 
    e.detalle
FROM 
    eventos_reporte.tmp_evento_error e
INNER JOIN 
    eventos_reporte.reporte_status rs 
    ON e.fk_status_reporte = rs.id
WHERE 
    rs.descripcion != 'RESUELTA'
    AND (e.codigo_seguimiento IS NULL OR e.codigo_seguimiento = '')
    AND e.fecha_creacion >= current_date - INTERVAL '5 days'
ORDER BY 
    e.fecha_creacion DESC;
`

const QUERY_GET_MUNICIPALITY_LOCATION = `
SELECT 
    tm.descripcion AS ciudad,
    td.nombre_departamento AS departamento
FROM public.tbl_municipios tm
INNER JOIN public.tbl_departamentos td 
    ON tm.tbl_departamentos_id = td.id_departamento
WHERE tm.id = $1;
`
