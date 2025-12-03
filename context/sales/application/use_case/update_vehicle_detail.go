package usecase

import (
	"log"

	entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
	irepo "ms-genexis-pos-operaciones/context/sales/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type UpdateVehicleDetail struct {
	Repository irepo.IUpdateVehicleDetailRepository
}

func (u *UpdateVehicleDetail) Execute(request *entities_sales.UpdateVehicleDetailRequest) (*entities_main.Response[entities_sales.UpdateVehicleDetailResult], error) {
	result, err := u.Repository.Update(request)
	if err != nil {
		log.Println("UpdateVehicleDetail usecase error:", err)
		return nil, err
	}
	return result, nil
}
