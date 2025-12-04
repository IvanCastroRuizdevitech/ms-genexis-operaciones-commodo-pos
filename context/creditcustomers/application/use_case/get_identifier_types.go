package use_case

import (
	"context"

	"ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	irepositories "ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetIdentifierTypes struct {
	Repository irepositories.IIdentifierTypesRepository
}

func (uc *GetIdentifierTypes) Execute(ctx context.Context) (*entities_main.Response[[]entities.IdentifierType], error) {
	return uc.Repository.GetAll(ctx)
}
