package service

import (
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/configuration/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetParametersClient struct {
	GetParameters iusecase.IGetParameters
}

func (configuration *GetParametersClient) Execute() (*entities_main.Response[entities.Config], error) {
	db_response, err := configuration.GetParameters.Execute()
	if err != nil {
		return nil, err
	}

	return db_response, nil
}
