package usecase

import (
	"log"

	entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
	irepo "ms-genexis-pos-operaciones/context/sales/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetDispenserDetails struct {
	Repository irepo.IGetDispenserDetailsRepository
}

func (u *GetDispenserDetails) Execute() (*entities_main.Response[[]entities_sales.DispenserDetail], error) {
	result, err := u.Repository.Get()
	if err != nil {
		log.Println("GetDispenserDetails usecase error:", err)
		return nil, err
	}
	return result, nil
}
