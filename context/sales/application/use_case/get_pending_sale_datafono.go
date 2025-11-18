package usecase

import (
    "log"
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    irepo "ms-genexis-pos-operaciones/context/sales/domain/ports/repositories"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetPendingSaleDatafono struct {
    Repository irepo.IGetPendingSaleDatafonoRepository
}

func (u *GetPendingSaleDatafono) Execute(transactionId int) (*entities_main.Response[entities_sales.PendingSaleDatafono], error) {
    result, err := u.Repository.Get(transactionId)
    if err != nil {
        log.Println("GetPendingSaleDatafono usecase error:", err)
        return nil, err
    }
    return result, nil
}

