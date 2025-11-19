package service

import (
    "ms-genexis-pos-operaciones/context/configuration/domain/entities"
    iusecase "ms-genexis-pos-operaciones/context/configuration/domain/ports/application/use_case"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetInitialConfigurationClient struct {
    GetInitialConfiguration iusecase.IGetInitialConfiguration
}

func (s *GetInitialConfigurationClient) Execute() (*entities_main.Response[entities.InitialConfiguration], error) {
    dbResponse, err := s.GetInitialConfiguration.Execute()
    if err != nil {
        return nil, err
    }
    return dbResponse, nil
}

