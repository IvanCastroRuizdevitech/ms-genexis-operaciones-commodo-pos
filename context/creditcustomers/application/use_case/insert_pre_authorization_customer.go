package use_case

import (
	entities_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	irepositories "ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type InsertPreAuthorizationCustomer struct {
	Repository irepositories.IInsertPreAuthorizationCustomerRepository
}

func (u *InsertPreAuthorizationCustomer) Execute(request *entities_creditcustomers.InsertPreAuthorizationCustomerRequest) (*entities_main.Response[entities_creditcustomers.InsertPreAuthorizationCustomerResult], error) {
	return u.Repository.Insert(request)
}
