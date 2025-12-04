package iusecase

import (
	"context"

	"ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IGetCreditCustomers interface {
	Execute(ctx context.Context) (*entities_main.Response[[]entities.CreditCustomer], error)
}
