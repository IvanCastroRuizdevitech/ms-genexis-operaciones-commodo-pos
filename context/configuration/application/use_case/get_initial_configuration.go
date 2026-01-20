package usecase

import (
	"log"

	entities_configuration "ms-genexis-pos-operaciones/context/configuration/domain/entities"
	irepo "ms-genexis-pos-operaciones/context/configuration/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetInitialConfiguration struct {
	Repository irepo.IGetInitialConfigurationRepository
}

func (u *GetInitialConfiguration) Execute() (*entities_main.Response[entities_configuration.InitialConfigurationDataResponse], error) {
	result, err := u.Repository.GetInitialConfiguration()
	if err != nil {
		log.Println("GetInitialConfiguration usecase error:", err)
		return nil, err
	}
	return result, nil
}
