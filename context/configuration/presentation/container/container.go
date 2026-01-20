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

var dbConn dbclient.DatabaseConnectionInterface

// REPOSITORIES DB
var GetInitialConfigurationRepository irepositories.IGetInitialConfigurationRepository

// USECASE
var GetInitialConfigurationUseCase iusecase.IGetInitialConfiguration

// SERVICE
var GetInitialConfigurationClient iservice.IGetInitialConfiguration

func resolveDB() dbclient.DatabaseConnectionInterface {
	if dbConn == nil {
		dbConn = presentation_container.ResolveDatabaseConnectionToLecWithPgx()
	}
	return dbConn
}

func buildGetInitialConfiguration() {
	if GetInitialConfigurationClient != nil {
		return
	}
	repo := &repositories.GetInitialConfigurationRepository{Connection: resolveDB()}
	GetInitialConfigurationRepository = repo
	GetInitialConfigurationUseCase = &usecase.GetInitialConfiguration{Repository: repo}
	GetInitialConfigurationClient = &service.GetInitialConfigurationClient{UseCase: GetInitialConfigurationUseCase}
}

func ResolveGetInitialConfigurationContainer() iservice.IGetInitialConfiguration {
	buildGetInitialConfiguration()
	return GetInitialConfigurationClient
}
