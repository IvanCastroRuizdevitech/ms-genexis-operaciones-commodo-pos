package irepositories

import (
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IAssignCustomerDataRepository interface {
    Assign(request *entities_sales.AssignCustomerDataRequest) (*entities_main.Response[entities_sales.AssignCustomerDataResult], error)
}

