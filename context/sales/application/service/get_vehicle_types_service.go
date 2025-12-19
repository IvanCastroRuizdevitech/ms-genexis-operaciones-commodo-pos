package service

import (
	entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/sales/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetVehicleTypesClient struct {
	UseCase iusecase.IGetVehicleTypes
}

func (s *GetVehicleTypesClient) Execute() (*entities_main.Response[entities_sales.VehicleTypesResponse], error) {
	return s.UseCase.Execute()
}
