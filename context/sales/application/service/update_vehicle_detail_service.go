package service

import (
	entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/sales/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type UpdateVehicleDetailClient struct {
	UseCase iusecase.IUpdateVehicleDetail
}

func (s *UpdateVehicleDetailClient) Execute(request *entities_sales.UpdateVehicleDetailRequest) (*entities_main.Response[entities_sales.UpdateVehicleDetailResult], error) {
	return s.UseCase.Execute(request)
}
