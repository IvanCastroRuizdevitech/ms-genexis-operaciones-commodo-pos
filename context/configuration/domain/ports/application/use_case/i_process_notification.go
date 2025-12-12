package iusecase

import (
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IProcessNotification interface {
	Execute(request *entities.ProcessNotificationRequest) (*entities_main.Response[entities.ProcessNotificationResult], error)
}
