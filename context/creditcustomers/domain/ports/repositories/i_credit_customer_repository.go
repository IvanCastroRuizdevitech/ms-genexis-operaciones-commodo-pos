package repositories

import (
	"context"

	"ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type ICreditCustomerRepository interface {
	GetAll(ctx context.Context) (*entities_main.Response[[]entities.CreditCustomer], error)
	UpdateLimit(ctx context.Context, request *entities.UpdateCreditCustomerLimitRequest) (*entities_main.Response[entities.UpdateCreditCustomerLimitResult], error)
}
