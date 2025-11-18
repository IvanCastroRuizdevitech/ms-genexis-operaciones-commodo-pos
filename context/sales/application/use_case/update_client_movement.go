package usecase

import (
    "log"
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    irepo "ms-genexis-pos-operaciones/context/sales/domain/ports/repositories"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type UpdateClientMovement struct {
    Repository irepo.IUpdateClientMovementRepository
}

func (u *UpdateClientMovement) Execute(request *entities_sales.UpdateClientMovementRequest) (*entities_main.Response[entities_sales.UpdateClientMovementResult], error) {
    result, err := u.Repository.Update(request)
    if err != nil {
        log.Println("UpdateClientMovement usecase error:", err)
        return nil, err
    }
    return result, nil
}

