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

// Repositories
var userRepository irepositories.IUserRepository

// Use cases
var getUsersUseCase iusecase.IGetUsers
var assignTagUseCase iusecase.IAssignTag
var assignTagTransmissionsUseCase iusecase.IAssignTagTransmissions

// Services
var usersService iservice.IUsersService
var assignTagService iservice.IAssignTag
var assignTagTransmissionsService iservice.IAssignTagTransmissions

func resolveDB() presentation_container.DatabaseConnectionInterface {
	return presentation_container.ResolveDatabaseConnectionToLecWithPgx()
}

func buildUsers() {
	if usersService != nil {
		return
	}
	userRepository = &repositories.UserRepository{
		Connection: resolveDB(),
	}
	getUsersUseCase = &usecase.GetUsers{
		Repository: userRepository,
	}
	usersService = &service.UsersService{
		GetUsers: getUsersUseCase,
	}
}

func buildAssignTag() {
	if assignTagService != nil {
		return
	}
	if userRepository == nil {
		buildUsers()
	}
	assignTagUseCase = &usecase.AssignTag{
		Repository: userRepository,
	}
	assignTagService = &service.AssignTagService{
		AssignTagUseCase: assignTagUseCase,
	}
}

func buildAssignTagTransmissions() {
	if assignTagTransmissionsService != nil {
		return
	}
	if userRepository == nil {
		buildUsers()
	}
	assignTagTransmissionsUseCase = &usecase.AssignTagTransmissions{
		Repository: userRepository,
	}
	assignTagTransmissionsService = &service.AssignTagTransmissionsService{
		AssignTagTransmissionsUseCase: assignTagTransmissionsUseCase,
	}
}

func ResolveUsersContainer() iservice.IUsersService {
	buildUsers()
	return usersService
}

func ResolveAssignTagContainer() iservice.IAssignTag {
	buildAssignTag()
	return assignTagService
}

func ResolveAssignTagTransmissionsContainer() iservice.IAssignTagTransmissions {
	buildAssignTagTransmissions()
	return assignTagTransmissionsService
}
