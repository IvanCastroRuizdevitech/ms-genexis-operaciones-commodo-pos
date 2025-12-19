package usecase

import (
	"log"

	entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
	irepo "ms-genexis-pos-operaciones/context/sales/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetVehicleTypes struct {
	Repository irepo.IGetVehicleTypesRepository
}

func (u *GetVehicleTypes) Execute() (*entities_main.Response[entities_sales.VehicleTypesResponse], error) {
	result, err := u.Repository.Get()
	if err != nil {
		log.Println("GetVehicleTypes usecase error:", err)
		return nil, err
	}
	return result, nil
}
