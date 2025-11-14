package usecase

import (
    "log"
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    irepo "ms-genexis-pos-operaciones/context/sales/domain/ports/repositories"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type CheckDatafonoCancellationsInProgress struct {
    Repository irepo.ICheckDatafonoCancellationsInProgressRepository
}

func (u *CheckDatafonoCancellationsInProgress) Execute(request *entities_sales.DatafonoCancellationsInProgressRequest) (*entities_main.Response[entities_sales.DatafonoCancellationsInProgress], error) {
    result, err := u.Repository.Check(request)
    if err != nil {
        log.Println("CheckDatafonoCancellationsInProgress usecase error:", err)
        return nil, err
    }
    return result, nil
}

