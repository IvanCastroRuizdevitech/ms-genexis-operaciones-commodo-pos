package irepositories

import (
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IProcessNotificationRepository interface {
	Process(request *entities.ProcessNotificationRequest) (*entities_main.Response[entities.ProcessNotificationResult], error)
}
