package usecase

import (
	"log"

	irepositories "ms-genexis-pos-operaciones/context/sales/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type UpsertActiveDispenserFaceTransaction struct {
	Repository irepositories.IUpsertActiveDispenserFaceTransactionRepository
}

func (u *UpsertActiveDispenserFaceTransaction) Execute(surtidor int, cara int, codigo string, grado int, proveedorID int, montoMaximo float64, cantidadMaxima float64, trama map[string]any, promotorID int) (*entities_main.Response[map[string]any], error) {
	result, err := u.Repository.Execute(surtidor, cara, codigo, grado, proveedorID, montoMaximo, cantidadMaxima, trama, promotorID)
	if err != nil {
		log.Println("UpsertActiveDispenserFaceTransaction usecase error:", err)
		return nil, err
	}
	return result, nil
}
