package service

import (
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/configuration/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetNotificationTypesClient struct {
	UseCase iusecase.IGetNotificationTypes
}

func (s *GetNotificationTypesClient) Execute() (*entities_main.Response[entities.NotificationTypesResponse], error) {
	return s.UseCase.Execute()
}
