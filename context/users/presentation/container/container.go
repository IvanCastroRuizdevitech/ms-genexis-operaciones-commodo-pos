package container_users

import (
	"ms-genexis-pos-operaciones/context/users/application/service"
	usecase "ms-genexis-pos-operaciones/context/users/application/use_case"
	iservice "ms-genexis-pos-operaciones/context/users/domain/ports/application/service"
	iusecase "ms-genexis-pos-operaciones/context/users/domain/ports/application/use_case"
	irepositories "ms-genexis-pos-operaciones/context/users/domain/ports/repositories"
	repositories "ms-genexis-pos-operaciones/context/users/infrastructure"
	presentation_container "ms-genexis-pos-operaciones/presentation/container"
)

var userRepository irepositories.IUserRepository
var getUsersUseCase iusecase.IGetUsers
var usersService iservice.IUsersService

func initializes() {
	userRepository = &repositories.UserRepository{
		Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx(),
	}
	getUsersUseCase = &usecase.GetUsers{
		Repository: userRepository,
	}
	usersService = &service.UsersService{
		GetUsers: getUsersUseCase,
	}
}

func ResolveUsersContainer() iservice.IUsersService {
	if usersService == nil {
		initializes()
	}
	return usersService
}
