package iusecase

import (
	"context"

	"ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type ICreatePreAuthorization interface {
	Execute(ctx context.Context, req *entities.PreAuthorizationRequest) (*entities_main.Response[bool], error)
}
