package service

import (
	"ms-genexis-pos-operaciones/context/users/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/users/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type AssignTagTransmissionsService struct {
	AssignTagTransmissionsUseCase iusecase.IAssignTagTransmissions
}

func (s *AssignTagTransmissionsService) Execute(request *entities.AssignTagTransmissionsRequest) (*entities_main.Response[entities.AssignTagTransmissionsResult], error) {
	return s.AssignTagTransmissionsUseCase.Execute(request)
}
