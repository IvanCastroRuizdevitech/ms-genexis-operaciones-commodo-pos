package iservice

import (
	entities_configuration "ms-genexis-pos-operaciones/context/configuration/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IGetInitialConfiguration interface {
	Execute() (*entities_main.Response[entities_configuration.InitialConfigurationDataResponse], error)
}
