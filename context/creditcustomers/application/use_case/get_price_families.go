package use_case

import (
	"context"

	"ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	irepositories "ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetPriceFamilies struct {
	Repository irepositories.IPriceFamiliesRepository
}

func (uc *GetPriceFamilies) Execute(ctx context.Context) (*entities_main.Response[[]entities.PriceFamily], error) {
	return uc.Repository.GetAll(ctx)
}
