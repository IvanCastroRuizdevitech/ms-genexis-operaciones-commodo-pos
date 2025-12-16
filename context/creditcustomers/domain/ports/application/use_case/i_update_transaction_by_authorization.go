package iusecase

import (
	entities_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IUpdateTransactionByAuthorization interface {
	Execute(request *entities_creditcustomers.UpdateTransactionByAuthorizationRequest) (*entities_main.Response[entities_creditcustomers.UpdateTransactionByAuthorizationResult], error)
}
