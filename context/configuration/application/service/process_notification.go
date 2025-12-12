package service

import (
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/configuration/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type ProcessNotificationClient struct {
	UseCase iusecase.IProcessNotification
}

func (s *ProcessNotificationClient) Execute(request *entities.ProcessNotificationRequest) (*entities_main.Response[entities.ProcessNotificationResult], error) {
	return s.UseCase.Execute(request)
}
