package usecase

import (
    "log"
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    irepo "ms-genexis-pos-operaciones/context/sales/domain/ports/repositories"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type CheckPendingSales struct {
    Repository irepo.ICheckPendingSalesRepository
}

func (u *CheckPendingSales) Execute(request *entities_sales.CheckPendingSalesRequest) (*entities_main.Response[[]entities_sales.PendingSale], error) {
    result, err := u.Repository.Check(request)
    if err != nil {
        log.Println("CheckPendingSales usecase error:", err)
        return nil, err
    }
    return result, nil
}

