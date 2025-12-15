package usecase

import (
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	irepositories "ms-genexis-pos-operaciones/context/configuration/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetNotificationTypes struct {
	Repository irepositories.IGetNotificationTypesRepository
}

func (u *GetNotificationTypes) Execute() (*entities_main.Response[entities.NotificationTypesResponse], error) {
	return u.Repository.GetNotificationTypes()
}
