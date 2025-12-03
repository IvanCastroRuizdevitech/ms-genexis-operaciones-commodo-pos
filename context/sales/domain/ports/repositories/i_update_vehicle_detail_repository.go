package irepositories

import (
	entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IUpdateVehicleDetailRepository interface {
	Update(request *entities_sales.UpdateVehicleDetailRequest) (*entities_main.Response[entities_sales.UpdateVehicleDetailResult], error)
}
