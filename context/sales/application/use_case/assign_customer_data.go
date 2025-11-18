package usecase

import (
    "log"
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    irepo "ms-genexis-pos-operaciones/context/sales/domain/ports/repositories"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type AssignCustomerData struct {
    Repository irepo.IAssignCustomerDataRepository
}

func (u *AssignCustomerData) Execute(request *entities_sales.AssignCustomerDataRequest) (*entities_main.Response[entities_sales.AssignCustomerDataResult], error) {
    result, err := u.Repository.Assign(request)
    if err != nil {
        log.Println("AssignCustomerData usecase error:", err)
        return nil, err
    }
    return result, nil
}

