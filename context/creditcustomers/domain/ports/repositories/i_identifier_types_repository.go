package repositories

import (
	"context"

	"ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IIdentifierTypesRepository interface {
	GetAll(ctx context.Context) (*entities_main.Response[[]entities.IdentifierType], error)
}
