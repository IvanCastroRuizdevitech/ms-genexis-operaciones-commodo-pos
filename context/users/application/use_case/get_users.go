package usecase

import (
	"log"
	"ms-genexis-pos-operaciones/context/users/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/users/domain/ports/application/use_case"
	irepositories "ms-genexis-pos-operaciones/context/users/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

var _ iusecase.IGetUsers = (*GetUsers)(nil)

type GetUsers struct {
	Repository irepositories.IUserRepository
}

func (u *GetUsers) Execute() (*entities_main.Response[[]entities.User], error) {
	result, err := u.Repository.GetAll()
	if err != nil {
		log.Println("[GetUsers][Execute]", err)
		return nil, err
	}
	return result, nil
}
