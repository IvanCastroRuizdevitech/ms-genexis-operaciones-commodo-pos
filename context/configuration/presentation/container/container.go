package container_configuration

import (
	"ms-genexis-pos-operaciones/context/configuration/application/service"
	usecase "ms-genexis-pos-operaciones/context/configuration/application/use_case"
	iservice "ms-genexis-pos-operaciones/context/configuration/domain/ports/application/service"
	iusecase "ms-genexis-pos-operaciones/context/configuration/domain/ports/application/use_case"
	irepositories "ms-genexis-pos-operaciones/context/configuration/domain/ports/repositories"
	repositories "ms-genexis-pos-operaciones/context/configuration/infrastructure"
	dbclient "ms-genexis-pos-operaciones/infrastructure/db/client"
	presentation_container "ms-genexis-pos-operaciones/presentation/container"
)

// REPOSITORIES DB
var GetParametersRepository irepositories.IGetParametersRepository
var GetPromoterDutyRepository irepositories.IGetPromoterDutyRepository
var GetInitialConfigurationRepository irepositories.IGetInitialConfigurationRepository

// REPOSITORIES HTTPP

// USECASE
var GetParametersUseCase iusecase.IGetParameters
var GetPromoterDutyUseCase iusecase.IGetPromoterDuty
var GetInitialConfigurationUseCase iusecase.IGetInitialConfiguration

// SERVICE
var GetParametersClient iservice.IGetParameters
var GetPromoterDutyClient iservice.IGetPromoterDuty
var GetInitialConfigurationClient iservice.IGetInitialConfiguration

func resolveDB() dbclient.DatabaseConnectionInterface {
	return presentation_container.ResolveDatabaseConnectionToLecWithPgx()
}

func buildGetParameters() {
	if GetParametersClient != nil {
		return
	}
	dbConn := resolveDB()
	GetParametersRepository = &repositories.GetParametersRepository{Connection: dbConn}
	GetParametersUseCase = &usecase.GetParameters{GetParameters: GetParametersRepository}
	GetParametersClient = &service.GetParametersClient{GetParameters: GetParametersUseCase}
}

func buildGetPromoterDuty() {
	if GetPromoterDutyClient != nil {
		return
	}
	dbConn := resolveDB()
	GetPromoterDutyRepository = &repositories.GetPromoterDutyRepository{Connection: dbConn}
	GetPromoterDutyUseCase = &usecase.GetPromoterDuty{GetPromoterDuty: GetPromoterDutyRepository}
	GetPromoterDutyClient = &service.GetPromoterDutyClient{GetPromoterDuty: GetPromoterDutyUseCase}
}

func buildGetInitialConfiguration() {
	if GetInitialConfigurationClient != nil {
		return
	}
	dbConn := resolveDB()
	GetInitialConfigurationRepository = &repositories.GetInitialConfigurationRepository{Connection: dbConn}
	GetInitialConfigurationUseCase = &usecase.GetInitialConfiguration{GetInitialConfiguration: GetInitialConfigurationRepository}
	GetInitialConfigurationClient = &service.GetInitialConfigurationClient{GetInitialConfiguration: GetInitialConfigurationUseCase}
}

func ResolveGetParametersContainer() iservice.IGetParameters {

	buildGetParameters()

	return GetParametersClient
}

func ResolveGetPromoterDutyContainer() iservice.IGetPromoterDuty {
	buildGetPromoterDuty()
	return GetPromoterDutyClient
}

func ResolveGetInitialConfigurationContainer() iservice.IGetInitialConfiguration {
	buildGetInitialConfiguration()
	return GetInitialConfigurationClient
}
