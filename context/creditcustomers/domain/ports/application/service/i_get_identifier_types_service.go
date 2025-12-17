package iservice

import (
	entities_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IGetIdentifierTypesService interface {
	Execute() (*entities_main.Response[[]entities_creditcustomers.IdentifierType], error)
}
