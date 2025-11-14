package usecase

import (
    "log"
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    irepo "ms-genexis-pos-operaciones/context/sales/domain/ports/repositories"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type CheckReadySales struct {
    Repository irepo.ICheckReadySalesRepository
}

func (u *CheckReadySales) Execute(request *entities_sales.CheckReadySalesRequest) (*entities_main.Response[[]entities_sales.PendingSale], error) {
    result, err := u.Repository.Check(request)
    if err != nil {
        log.Println("CheckReadySales usecase error:", err)
        return nil, err
    }
    return result, nil
}

