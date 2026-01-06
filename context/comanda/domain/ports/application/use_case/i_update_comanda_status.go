package iusecase

import (
	entities_comanda "ms-genexis-pos-operaciones/context/comanda/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IUpdateComandaStatus interface {
	Execute(request *entities_comanda.UpdateComandaStatusRequest) (*entities_main.Response[entities_comanda.UpdateComandaStatusResult], error)
}
