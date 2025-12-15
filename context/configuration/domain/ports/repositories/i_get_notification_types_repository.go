package irepositories

import (
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IGetNotificationTypesRepository interface {
	GetNotificationTypes() (*entities_main.Response[entities.NotificationTypesResponse], error)
}
