package service

import (
	"ms-genexis-pos-operaciones/context/users/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/users/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type UsersService struct {
	GetUsers iusecase.IGetUsers
}

func (s *UsersService) Execute() (*entities_main.Response[[]entities.User], error) {
	return s.GetUsers.Execute()
}
