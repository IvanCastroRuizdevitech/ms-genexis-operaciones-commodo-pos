package use_case

import (
	"context"

	"ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	"ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetCreditCustomers struct {
	Repository repositories.ICreditCustomerRepository
}

func (uc *GetCreditCustomers) Execute(ctx context.Context) (*entities_main.Response[[]entities.CreditCustomer], error) {
	return uc.Repository.GetAll(ctx)
}
