package irepositories

import (
	"ms-genexis-pos-operaciones/context/home/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IGetMunicipalityLocationRepository interface {
	GetByID(id int) (*entities_main.Response[entities.MunicipalityLocation], error)
}
