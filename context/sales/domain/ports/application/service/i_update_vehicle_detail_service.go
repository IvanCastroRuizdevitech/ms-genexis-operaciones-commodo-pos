package iservice

import (
	entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IUpdateVehicleDetail interface {
	Execute(request *entities_sales.UpdateVehicleDetailRequest) (*entities_main.Response[entities_sales.UpdateVehicleDetailResult], error)
}
