package iservice

import (
	entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IGetVehicleTypes interface {
	Execute() (*entities_main.Response[entities_sales.VehicleTypesResponse], error)
}
