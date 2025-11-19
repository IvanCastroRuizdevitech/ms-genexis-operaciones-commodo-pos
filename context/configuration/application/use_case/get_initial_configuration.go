package usecase

import (
    "log"
    "ms-genexis-pos-operaciones/context/configuration/domain/entities"
    irepositories "ms-genexis-pos-operaciones/context/configuration/domain/ports/repositories"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetInitialConfiguration struct {
    GetInitialConfiguration irepositories.IGetInitialConfigurationRepository
}

func (u *GetInitialConfiguration) Execute() (*entities_main.Response[entities.InitialConfiguration], error) {
    result, err := u.GetInitialConfiguration.GetInitialConfiguration()
    if err != nil {
        log.Println("GetInitialConfiguration: ", err)
        return nil, err
    }
    return result, nil
}

