package irepositories

import (
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IGetParametersRepository interface {
	GetParameters() (*entities_main.Response[entities.Config], error)
}
