package iusecase

import (
	entities_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IInsertPreAuthorizationCustomer interface {
	Execute(request *entities_creditcustomers.InsertPreAuthorizationCustomerRequest) (*entities_main.Response[entities_creditcustomers.InsertPreAuthorizationCustomerResult], error)
}
