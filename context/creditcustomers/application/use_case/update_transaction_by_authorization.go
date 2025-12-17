package use_case

import (
	entities_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	irepositories "ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type UpdateTransactionByAuthorization struct {
	Repository irepositories.IUpdateTransactionByAuthorizationRepository
}

func (u *UpdateTransactionByAuthorization) Execute(request *entities_creditcustomers.UpdateTransactionByAuthorizationRequest) (*entities_main.Response[entities_creditcustomers.UpdateTransactionByAuthorizationResult], error) {
	return u.Repository.Update(request)
}
