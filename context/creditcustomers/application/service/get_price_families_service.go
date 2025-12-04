package service

import (
	"context"

	"ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetPriceFamiliesService struct {
	UseCase iusecase.IGetPriceFamilies
}

func (s *GetPriceFamiliesService) Execute(ctx context.Context) (*entities_main.Response[[]entities.PriceFamily], error) {
	return s.UseCase.Execute(ctx)
}
