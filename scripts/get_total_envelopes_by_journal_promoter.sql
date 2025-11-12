CREATE OR REPLACE FUNCTION public.get_total_envelopes_by_journal_promoter(
    i_journal_id int,
    i_promoter_id int
)
RETURNS JSONB
LANGUAGE plpgsql
-- =========================================================================================
-- Nombre:          get_total_envelopes_by_journal_promoter
-- Descripción:     Obtiene el total de ventas por sobres (efectivo) para una jornada y un promotor.
-- Historia de Usuario:  Migración funcionalidad Interfaz a PL
-- Autor:           Carlos De Las Salas Ospino
-- Sprint:          N/A
-- Fecha Creación:  12/11/2025
-- =========================================================================================
AS $function$
DECLARE
    c_PAYMENT_METHOD_CASH constant INT := 1;
    v_total NUMERIC := 0;
    v_result JSONB;
BEGIN
    SELECT
        COALESCE(SUM(cmmp.valor_total), 0)
    INTO v_total
    FROM ct_movimientos cm
    INNER JOIN ct_movimientos_medios_pagos cmmp 
        ON cm.id = cmmp.ct_movimientos_id
    WHERE 
        cmmp.ct_medios_pagos_id = c_PAYMENT_METHOD_CASH
        AND cm.jornadas_id = i_journal_id
        AND cm.responsables_id = i_promoter_id
        AND cm.tipo NOT IN ('016', '014');

    v_result := JSONB_BUILD_OBJECT(
        'status', 200,
        'message', 'Total envelopes retrieved successfully',
        'process_date', TO_CHAR(NOW(), 'YYYY-MM-DD HH24:MI:SS'),
        'data', JSONB_BUILD_OBJECT(
            'total', v_total
        )
    );

    RETURN v_result;

EXCEPTION
    WHEN OTHERS THEN
        RETURN JSONB_BUILD_OBJECT(
            'status', 500,
            'message', 'Error retrieving total envelopes: ' || SQLERRM,
            'process_date', TO_CHAR(NOW(), 'YYYY-MM-DD HH24:MI:SS'),
            'data', JSONB_BUILD_OBJECT(
                'total', 0
            )
        );
END;
$function$;
