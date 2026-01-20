package service

import (
	entities_configuration "ms-genexis-pos-operaciones/context/configuration/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/configuration/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetInitialConfigurationClient struct {
	UseCase iusecase.IGetInitialConfiguration
}

func (s *GetInitialConfigurationClient) Execute() (*entities_main.Response[entities_configuration.InitialConfigurationDataResponse], error) {
	return s.UseCase.Execute()
}
