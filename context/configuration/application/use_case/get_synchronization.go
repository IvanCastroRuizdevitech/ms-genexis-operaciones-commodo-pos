package usecase

import (
	"log"

	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	irepositories "ms-genexis-pos-operaciones/context/configuration/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetSynchronization struct {
	Repository irepositories.IGetSynchronizationRepository
}

func (u *GetSynchronization) Execute(idSincronizacion int, fechaInicio string, fechaFin string) (*entities_main.Response[[]entities.SynchronizationDetail], error) {
	result, err := u.Repository.Get(idSincronizacion, fechaInicio, fechaFin)
	if err != nil {
		log.Println("[GetSynchronization][Execute]", err)
		return nil, err
	}
	return result, nil
}
