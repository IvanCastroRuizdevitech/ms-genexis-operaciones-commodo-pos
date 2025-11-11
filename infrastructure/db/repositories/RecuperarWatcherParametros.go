package repositorios_infraestruture

import (
	domain_adapters_clients_db "ms-genexis-pos-operaciones/domain/adapters/clients/db"
	"ms-genexis-pos-operaciones/domain/entities"
)

type RecuperarWatcherParametors struct {
	Cliente domain_adapters_clients_db.IClientDB
}

func (RWP *RecuperarWatcherParametors) Consultar(codigo string) (*entities.ParametersWatcher, error) {
	args := []any{codigo}
	respuesta, err := RWP.Cliente.Select("SELECT x.* FROM public.wacher_parametros x WHERE codigo = $1", args)

	if err != nil {
		return nil, err
	}

	parametro := &entities.ParametersWatcher{}
	for _, valor := range respuesta {
		parametro.Id = valor[0].(int64)
		parametro.Codigo = valor[1].(string)
		parametro.Tipo = valor[2].(int32)
		parametro.Valor = valor[3].(string)

	}

	return parametro, nil
}
