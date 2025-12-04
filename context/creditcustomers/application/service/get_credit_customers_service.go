package service

import (
	"context"

	"ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetCreditCustomersService struct {
	UseCase iusecase.IGetCreditCustomers
}

func (s *GetCreditCustomersService) Execute(ctx context.Context) (*entities_main.Response[[]entities.CreditCustomer], error) {
	return s.UseCase.Execute(ctx)
}
