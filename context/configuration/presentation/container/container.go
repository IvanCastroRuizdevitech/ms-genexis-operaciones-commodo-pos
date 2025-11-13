package container_configuration

import (
	"ms-genexis-pos-operaciones/context/configuration/application/service"
	usecase "ms-genexis-pos-operaciones/context/configuration/application/use_case"
	iservice "ms-genexis-pos-operaciones/context/configuration/domain/ports/application/service"
	iusecase "ms-genexis-pos-operaciones/context/configuration/domain/ports/application/use_case"
	irepositories "ms-genexis-pos-operaciones/context/configuration/domain/ports/repositories"
	repositories "ms-genexis-pos-operaciones/context/configuration/infrastructure"
	presentation_container "ms-genexis-pos-operaciones/presentation/container"
)

// REPOSITORIES DB
var GetParametersRepository irepositories.IGetParametersRepository

// REPOSITORIES HTTPP

// USECASE
var GetParametersUseCase iusecase.IGetParameters

// SERVICE
var GetParametersClient iservice.IGetParameters

func initializes() {
	GetParametersRepository = &repositories.GetParametersRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}
	GetParametersUseCase = &usecase.GetParameters{GetParameters: GetParametersRepository}
	GetParametersClient = &service.GetParametersClient{GetParameters: GetParametersUseCase}
}

func ResolveGetParametersContainer() iservice.IGetParameters {

	if GetParametersClient == nil {
		initializes()
	}

	return GetParametersClient
}
