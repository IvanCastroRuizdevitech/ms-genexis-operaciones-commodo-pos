package service

import (
	entities_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetDispenserDetailsByFamiliesService struct {
	UseCase iusecase.IGetDispenserDetailsByFamilies
}

func (s *GetDispenserDetailsByFamiliesService) Execute(request *entities_creditcustomers.DispenserDetailsByFamiliesRequest) (*entities_main.Response[entities_creditcustomers.DispenserDetailsFunctionResponse], error) {
	return s.UseCase.Execute(request)
}
