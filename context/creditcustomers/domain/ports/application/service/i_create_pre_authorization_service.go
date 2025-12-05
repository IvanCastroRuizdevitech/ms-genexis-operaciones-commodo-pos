package iservice

import (
	"context"

	"ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type ICreatePreAuthorizationService interface {
	Execute(ctx context.Context, req *entities.PreAuthorizationRequest) (*entities_main.Response[bool], error)
}
