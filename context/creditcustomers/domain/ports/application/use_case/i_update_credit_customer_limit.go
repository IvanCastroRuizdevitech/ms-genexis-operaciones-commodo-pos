package iusecase

import (
	"context"

	"ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IUpdateCreditCustomerLimit interface {
	Execute(ctx context.Context, request *entities.UpdateCreditCustomerLimitRequest) (*entities_main.Response[entities.UpdateCreditCustomerLimitResult], error)
}
