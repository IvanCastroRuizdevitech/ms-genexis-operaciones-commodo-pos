package usecase

import (
	"log"
	"ms-genexis-pos-operaciones/context/users/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/users/domain/ports/application/use_case"
	irepositories "ms-genexis-pos-operaciones/context/users/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

var _ iusecase.IAssignTagTransmissions = (*AssignTagTransmissions)(nil)

type AssignTagTransmissions struct {
	Repository irepositories.IUserRepository
}

func (u *AssignTagTransmissions) Execute(request *entities.AssignTagTransmissionsRequest) (*entities_main.Response[entities.AssignTagTransmissionsResult], error) {
	response, err := u.Repository.GenerateAssignTagTransmissions(request)
	if err != nil {
		log.Println("[AssignTagTransmissions][Execute]", err)
		return nil, err
	}
	return response, nil
}
