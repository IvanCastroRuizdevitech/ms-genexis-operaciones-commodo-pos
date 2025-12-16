package irepositories

import (
	entities_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IUpdateTransactionByAuthorizationRepository interface {
	Update(request *entities_creditcustomers.UpdateTransactionByAuthorizationRequest) (*entities_main.Response[entities_creditcustomers.UpdateTransactionByAuthorizationResult], error)
}
