package constants

const QueryGetUsers = `
SELECT * FROM public.fnc_obtener_personas_filtradas();
`

const QueryClearTag = `
UPDATE personas
SET tag = NULL
WHERE tag = $1;
`

const QueryAssignTag = `
UPDATE personas
SET tag = $1
WHERE identificacion = $2
RETURNING id;
`

const QueryGenerateAssignTagTransmissions = `
SELECT public.fnc_generar_transmisiones_asignar_tag($1, $2, $3);
`
