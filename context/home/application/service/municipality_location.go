package service

import (
	"ms-genexis-pos-operaciones/context/home/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/home/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type MunicipalityLocationClient struct {
	GetMunicipalityLocation iusecase.IGetMunicipalityLocation
}

func (s *MunicipalityLocationClient) Execute(id int) (*entities_main.Response[entities.MunicipalityLocation], error) {
	return s.GetMunicipalityLocation.Execute(id)
}
