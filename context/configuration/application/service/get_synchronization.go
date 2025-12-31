package service

import (
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/configuration/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetSynchronizationClient struct {
	GetSynchronization iusecase.IGetSynchronization
}

func (s *GetSynchronizationClient) Execute(idSincronizacion int, fechaInicio string, fechaFin string) (*entities_main.Response[[]entities.SynchronizationDetail], error) {
	return s.GetSynchronization.Execute(idSincronizacion, fechaInicio, fechaFin)
}
