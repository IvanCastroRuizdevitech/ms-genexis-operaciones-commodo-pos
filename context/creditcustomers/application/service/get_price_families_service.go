package service

import (
	entities_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetPriceFamiliesService struct {
	UseCase iusecase.IGetPriceFamilies
}

func (s *GetPriceFamiliesService) Execute() (*entities_main.Response[[]entities_creditcustomers.PriceFamily], error) {
	return s.UseCase.Execute()
}
