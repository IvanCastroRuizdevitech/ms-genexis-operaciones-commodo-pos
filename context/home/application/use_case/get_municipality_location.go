package usecase

import (
	"log"
	"ms-genexis-pos-operaciones/context/home/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/home/domain/ports/application/use_case"
	irepositories "ms-genexis-pos-operaciones/context/home/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

var _ iusecase.IGetMunicipalityLocation = (*GetMunicipalityLocation)(nil)

type GetMunicipalityLocation struct {
	Repository irepositories.IGetMunicipalityLocationRepository
}

func (u *GetMunicipalityLocation) Execute(id int) (*entities_main.Response[entities.MunicipalityLocation], error) {
	result, err := u.Repository.GetByID(id)
	if err != nil {
		log.Println("[GetMunicipalityLocation][Execute]", err)
		return nil, err
	}
	return result, nil
}
