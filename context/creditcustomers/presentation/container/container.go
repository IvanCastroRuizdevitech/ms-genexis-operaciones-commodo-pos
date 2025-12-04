package container_creditcustomers

import (
	app_service "ms-genexis-pos-operaciones/context/creditcustomers/application/service"
	usecase "ms-genexis-pos-operaciones/context/creditcustomers/application/use_case"
	iservice "ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/application/service"
	iusecase "ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/application/use_case"
	irepositories "ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/repositories"
	repositories "ms-genexis-pos-operaciones/context/creditcustomers/infrastructure"
	dbclient "ms-genexis-pos-operaciones/infrastructure/db/client"
	presentation_container "ms-genexis-pos-operaciones/presentation/container"
)

var dbConn dbclient.DatabaseConnectionInterface

var IdentifierTypesRepository irepositories.IIdentifierTypesRepository
var PriceFamiliesRepository irepositories.IPriceFamiliesRepository

var GetIdentifierTypesUseCase iusecase.IGetIdentifierTypes
var GetPriceFamiliesUseCase iusecase.IGetPriceFamilies

var GetIdentifierTypesClient iservice.IGetIdentifierTypesService
var GetPriceFamiliesClient iservice.IGetPriceFamiliesService

func resolveDB() dbclient.DatabaseConnectionInterface {
	if dbConn == nil {
		dbConn = presentation_container.ResolveDatabaseConnectionToLecWithPgx()
	}
	return dbConn
}

func buildGetIdentifierTypes() {
	if GetIdentifierTypesClient != nil {
		return
	}
	repo := &repositories.IdentifierTypesRepository{Connection: resolveDB()}
	IdentifierTypesRepository = repo
	GetIdentifierTypesUseCase = &usecase.GetIdentifierTypes{Repository: repo}
	GetIdentifierTypesClient = &app_service.GetIdentifierTypesService{UseCase: GetIdentifierTypesUseCase}
}

func ResolveGetIdentifierTypesContainer() iservice.IGetIdentifierTypesService {
	buildGetIdentifierTypes()
	return GetIdentifierTypesClient
}

func buildGetPriceFamilies() {
	if GetPriceFamiliesClient != nil {
		return
	}
	repo := &repositories.PriceFamiliesRepository{Connection: resolveDB()}
	PriceFamiliesRepository = repo
	GetPriceFamiliesUseCase = &usecase.GetPriceFamilies{Repository: repo}
	GetPriceFamiliesClient = &app_service.GetPriceFamiliesService{UseCase: GetPriceFamiliesUseCase}
}

func ResolveGetPriceFamiliesContainer() iservice.IGetPriceFamiliesService {
	buildGetPriceFamilies()
	return GetPriceFamiliesClient
}
