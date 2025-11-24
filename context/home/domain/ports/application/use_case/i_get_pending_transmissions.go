package iusecase

import (
	"ms-genexis-pos-operaciones/context/home/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IGetPendingTransmissions interface {
	Execute() (*entities_main.Response[[]entities.PendingTransmission], error)
}
