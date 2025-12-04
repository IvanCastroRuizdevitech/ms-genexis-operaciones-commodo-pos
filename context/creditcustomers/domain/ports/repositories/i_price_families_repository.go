package repositories

import (
	"context"

	"ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IPriceFamiliesRepository interface {
	GetAll(ctx context.Context) (*entities_main.Response[[]entities.PriceFamily], error)
}
