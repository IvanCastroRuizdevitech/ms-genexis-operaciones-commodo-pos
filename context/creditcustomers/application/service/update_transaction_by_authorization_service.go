package service

import (
	entities_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type UpdateTransactionByAuthorizationService struct {
	UseCase iusecase.IUpdateTransactionByAuthorization
}

func (s *UpdateTransactionByAuthorizationService) Execute(request *entities_creditcustomers.UpdateTransactionByAuthorizationRequest) (*entities_main.Response[entities_creditcustomers.UpdateTransactionByAuthorizationResult], error) {
	return s.UseCase.Execute(request)
}
