package repositories

import (
	"context"

	"ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
)

type IPreAuthorizationRepository interface {
	Create(ctx context.Context, req *entities.PreAuthorizationRequest) (bool, error)
}
