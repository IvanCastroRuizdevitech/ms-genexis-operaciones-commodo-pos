package service

import (
	entities_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type InsertPreAuthorizationCustomerService struct {
	UseCase iusecase.IInsertPreAuthorizationCustomer
}

func (s *InsertPreAuthorizationCustomerService) Execute(request *entities_creditcustomers.InsertPreAuthorizationCustomerRequest) (*entities_main.Response[entities_creditcustomers.InsertPreAuthorizationCustomerResult], error) {
	return s.UseCase.Execute(request)
}
