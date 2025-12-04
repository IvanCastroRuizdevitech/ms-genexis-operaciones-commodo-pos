package service

import (
	"context"

	"ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetIdentifierTypesService struct {
	UseCase iusecase.IGetIdentifierTypes
}

func (s *GetIdentifierTypesService) Execute(ctx context.Context) (*entities_main.Response[[]entities.IdentifierType], error) {
	return s.UseCase.Execute(ctx)
}
