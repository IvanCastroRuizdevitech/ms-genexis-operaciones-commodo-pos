package usecase

import (
    "log"
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    irepo "ms-genexis-pos-operaciones/context/sales/domain/ports/repositories"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type UpdateMovementState struct {
    Repository irepo.IUpdateMovementStateRepository
}

func (u *UpdateMovementState) Execute(movementId int, request *entities_sales.UpdateMovementStateRequest) (*entities_main.Response[entities_sales.UpdateMovementStateResult], error) {
    result, err := u.Repository.Update(movementId, request)
    if err != nil {
        log.Println("UpdateMovementState usecase error:", err)
        return nil, err
    }
    return result, nil
}

