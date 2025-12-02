package usecase

import (
	"log"

	irepositories "ms-genexis-pos-operaciones/context/sales/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetTransactionsByDispenserAndFace struct {
	Repository irepositories.IGetTransactionsByDispenserAndFaceRepository
}

func (u *GetTransactionsByDispenserAndFace) Execute(dispenserID int, face int) (*entities_main.Response[[]map[string]interface{}], error) {
	result, err := u.Repository.Get(dispenserID, face)
	if err != nil {
		log.Println("GetTransactionsByDispenserAndFace usecase error:", err)
		return nil, err
	}
	return result, nil
}
