package service

import (
	"context"

	"ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type UpdateCreditCustomerLimitService struct {
	UseCase iusecase.IUpdateCreditCustomerLimit
}

func (s *UpdateCreditCustomerLimitService) Execute(ctx context.Context, request *entities.UpdateCreditCustomerLimitRequest) (*entities_main.Response[entities.UpdateCreditCustomerLimitResult], error) {
	return s.UseCase.Execute(ctx, request)
}
