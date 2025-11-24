package usecase

import (
	"log"
	"ms-genexis-pos-operaciones/context/home/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/home/domain/ports/application/use_case"
	irepositories "ms-genexis-pos-operaciones/context/home/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

var _ iusecase.IGetPendingTransmissions = (*GetPendingTransmissions)(nil)

type GetPendingTransmissions struct {
	Repository irepositories.IGetPendingTransmissionsRepository
}

func (u *GetPendingTransmissions) Execute() (*entities_main.Response[[]entities.PendingTransmission], error) {
	result, err := u.Repository.GetAll()
	if err != nil {
		log.Println("[GetPendingTransmissions][Execute]", err)
		return nil, err
	}
	return result, nil
}
