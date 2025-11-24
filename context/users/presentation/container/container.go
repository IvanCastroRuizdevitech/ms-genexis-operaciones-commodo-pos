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
var assignTagUseCase iusecase.IAssignTag
var assignTagService iservice.IAssignTag
var assignTagTransmissionsUseCase iusecase.IAssignTagTransmissions
var assignTagTransmissionsService iservice.IAssignTagTransmissions

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

	assignTagUseCase = &usecase.AssignTag{
		Repository: userRepository,
	}
	assignTagService = &service.AssignTagService{
		AssignTagUseCase: assignTagUseCase,
	}

	assignTagTransmissionsUseCase = &usecase.AssignTagTransmissions{
		Repository: userRepository,
	}
	assignTagTransmissionsService = &service.AssignTagTransmissionsService{
		AssignTagTransmissionsUseCase: assignTagTransmissionsUseCase,
	}
}

func ResolveUsersContainer() iservice.IUsersService {
	if usersService == nil {
		initializes()
	}
	return usersService
}

func ResolveAssignTagContainer() iservice.IAssignTag {
	if assignTagService == nil {
		initializes()
	}
	return assignTagService
}

func ResolveAssignTagTransmissionsContainer() iservice.IAssignTagTransmissions {
	if assignTagTransmissionsService == nil {
		initializes()
	}
	return assignTagTransmissionsService
}
