package service

import (
	"ms-genexis-pos-operaciones/context/home/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/home/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type LoadErrorNotificationClient struct {
	LoadErrorNotification iusecase.ILoadErrorNotification
}

func (s *LoadErrorNotificationClient) Execute() (*entities_main.Response[[]entities.ErrorNotification], error) {
	return s.LoadErrorNotification.Execute()
}
