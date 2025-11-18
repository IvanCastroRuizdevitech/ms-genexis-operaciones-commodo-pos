package usecase

import (
    "log"
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    irepo "ms-genexis-pos-operaciones/context/sales/domain/ports/repositories"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type UpdatePaymentMethods struct {
    Repository irepo.IUpdatePaymentMethodsRepository
}

func (u *UpdatePaymentMethods) Execute(request *entities_sales.UpdatePaymentMethodsRequest) (*entities_main.Response[entities_sales.UpdatePaymentMethodsResult], error) {
    result, err := u.Repository.Update(request)
    if err != nil {
        log.Println("UpdatePaymentMethods usecase error:", err)
        return nil, err
    }
    return result, nil
}

