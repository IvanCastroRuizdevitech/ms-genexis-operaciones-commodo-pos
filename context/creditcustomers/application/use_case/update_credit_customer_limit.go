package use_case

import (
	"context"

	"ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	"ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type UpdateCreditCustomerLimit struct {
	Repository repositories.ICreditCustomerRepository
}

func (uc *UpdateCreditCustomerLimit) Execute(ctx context.Context, request *entities.UpdateCreditCustomerLimitRequest) (*entities_main.Response[entities.UpdateCreditCustomerLimitResult], error) {
	return uc.Repository.UpdateLimit(ctx, request)
}
