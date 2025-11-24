package usecase

import (
	"log"
	"ms-genexis-pos-operaciones/context/users/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/users/domain/ports/application/use_case"
	irepositories "ms-genexis-pos-operaciones/context/users/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

var _ iusecase.IAssignTag = (*AssignTag)(nil)

type AssignTag struct {
	Repository irepositories.IUserRepository
}

func (u *AssignTag) Execute(request *entities.AssignTagRequest) (*entities_main.Response[entities.AssignTagResult], error) {
	response, err := u.Repository.AssignTag(request)
	if err != nil {
		log.Println("[AssignTag][Execute]", err)
		return nil, err
	}
	return response, nil
}
