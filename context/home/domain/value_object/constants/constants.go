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

const QUERY_GET_PENDING_TRANSMISSIONS = `
SELECT public.fnc_obtener_transmisiones_pendientes();
`
