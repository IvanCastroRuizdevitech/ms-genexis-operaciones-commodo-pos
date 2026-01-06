create or replace function restaurante.fnc_actualizar_estado_comanda(
	p_comanda_id int8,
	p_nuevo_estado_id int2
)
returns json
language plpgsql
as $$
declare
	updated_rows int;
begin
	update restaurante.tbl_comandas_pedidos
	set estado_id = p_nuevo_estado_id
	where id = p_comanda_id;

	get diagnostics updated_rows = row_count;

	if updated_rows > 0 then
		return json_build_object(
			"success", true,
			"message", "Estado de comanda actualizado",
			"status", 200
		);
	end if;

	return json_build_object(
		"success", false,
		"message", "Comanda no encontrada",
		"status", 404
	);
end;
$$;
