package usecase

import (
	"log"
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	irepositories "ms-genexis-pos-operaciones/context/configuration/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetParameters struct {
	GetParameters irepositories.IGetParametersRepository
}

func (configuration *GetParameters) Execute() (*entities_main.Response[entities.Config], error) {

	result, err := configuration.GetParameters.GetParameters()

	if err != nil {
		log.Println("GetParameters: ", err)
		return nil, err
	}

	return result, nil
}
