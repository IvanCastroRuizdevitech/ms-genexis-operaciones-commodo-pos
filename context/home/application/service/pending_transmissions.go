package service

import (
	"ms-genexis-pos-operaciones/context/home/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/home/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type PendingTransmissionsClient struct {
	GetPendingTransmissions iusecase.IGetPendingTransmissions
}

func (s *PendingTransmissionsClient) Execute() (*entities_main.Response[[]entities.PendingTransmission], error) {
	return s.GetPendingTransmissions.Execute()
}
