package irepositories

import (
	"ms-genexis-pos-operaciones/context/home/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IGetPendingTransmissionsRepository interface {
	GetAll() (*entities_main.Response[[]entities.PendingTransmission], error)
}
