package iusecase

import (
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IGetSynchronization interface {
	Execute(idSincronizacion int, fechaInicio string, fechaFin string) (*entities_main.Response[[]entities.SynchronizationDetail], error)
}
