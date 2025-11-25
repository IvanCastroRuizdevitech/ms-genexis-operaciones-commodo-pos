package service

import (
	"ms-genexis-pos-operaciones/context/users/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/users/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type AssignTagService struct {
	AssignTagUseCase iusecase.IAssignTag
}

func (s *AssignTagService) Execute(request *entities.AssignTagRequest) (*entities_main.Response[entities.AssignTagResult], error) {
	return s.AssignTagUseCase.Execute(request)
}
