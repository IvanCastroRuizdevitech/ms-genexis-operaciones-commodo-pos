package service

import (
	entities_comanda "ms-genexis-pos-operaciones/context/comanda/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/comanda/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type UpdateComandaStatusClient struct {
	UseCase iusecase.IUpdateComandaStatus
}

func (s *UpdateComandaStatusClient) Execute(request *entities_comanda.UpdateComandaStatusRequest) (*entities_main.Response[entities_comanda.UpdateComandaStatusResult], error) {
	return s.UseCase.Execute(request)
}
