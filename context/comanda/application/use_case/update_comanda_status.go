package usecase

import (
	"log"

	entities_comanda "ms-genexis-pos-operaciones/context/comanda/domain/entities"
	irepo "ms-genexis-pos-operaciones/context/comanda/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type UpdateComandaStatus struct {
	Repository irepo.IUpdateComandaStatusRepository
}

func (u *UpdateComandaStatus) Execute(request *entities_comanda.UpdateComandaStatusRequest) (*entities_main.Response[entities_comanda.UpdateComandaStatusResult], error) {
	result, err := u.Repository.Update(request)
	if err != nil {
		log.Println("UpdateComandaStatus usecase error:", err)
		return nil, err
	}
	return result, nil
}
