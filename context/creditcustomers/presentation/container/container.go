package container_creditcustomers

import (
	"ms-genexis-pos-operaciones/context/creditcustomers/application/service"
	usecase "ms-genexis-pos-operaciones/context/creditcustomers/application/use_case"
	iservice "ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/application/service"
	iusecase "ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/application/use_case"
	irepositories "ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/repositories"
	repositories "ms-genexis-pos-operaciones/context/creditcustomers/infrastructure"
	presentation_container "ms-genexis-pos-operaciones/presentation/container"
)

var GetDispenserDetailsRepository irepositories.IGetDispenserDetailsRepository
var GetDispenserDetailsUseCase iusecase.IGetDispenserDetails
var GetDispenserDetailsClient iservice.IGetDispenserDetailsService
var GetIdentifierTypesRepository irepositories.IGetIdentifierTypesRepository
var GetIdentifierTypesUseCase iusecase.IGetIdentifierTypes
var GetIdentifierTypesClient iservice.IGetIdentifierTypesService
var GetPriceFamiliesRepository irepositories.IGetPriceFamiliesRepository
var GetPriceFamiliesUseCase iusecase.IGetPriceFamilies
var GetPriceFamiliesClient iservice.IGetPriceFamiliesService
var GetDispenserDetailsByFamiliesRepository irepositories.IGetDispenserDetailsByFamiliesRepository
var GetDispenserDetailsByFamiliesUseCase iusecase.IGetDispenserDetailsByFamilies
var GetDispenserDetailsByFamiliesClient iservice.IGetDispenserDetailsByFamiliesService
var UpdateTransactionByAuthorizationRepository irepositories.IUpdateTransactionByAuthorizationRepository
var UpdateTransactionByAuthorizationUseCase iusecase.IUpdateTransactionByAuthorization
var UpdateTransactionByAuthorizationClient iservice.IUpdateTransactionByAuthorizationService

func buildGetDispenserDetails() {
	if GetDispenserDetailsClient != nil {
		return
	}
	GetDispenserDetailsRepository = &repositories.GetDispenserDetailsRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}
	GetDispenserDetailsUseCase = &usecase.GetDispenserDetails{Repository: GetDispenserDetailsRepository}
	GetDispenserDetailsClient = &service.GetDispenserDetailsService{UseCase: GetDispenserDetailsUseCase}
}

func buildGetDispenserDetailsByFamilies() {
	if GetDispenserDetailsByFamiliesClient != nil {
		return
	}
	GetDispenserDetailsByFamiliesRepository = &repositories.GetDispenserDetailsByFamiliesRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}
	GetDispenserDetailsByFamiliesUseCase = &usecase.GetDispenserDetailsByFamilies{Repository: GetDispenserDetailsByFamiliesRepository}
	GetDispenserDetailsByFamiliesClient = &service.GetDispenserDetailsByFamiliesService{UseCase: GetDispenserDetailsByFamiliesUseCase}
}

func buildGetIdentifierTypes() {
	if GetIdentifierTypesClient != nil {
		return
	}
	GetIdentifierTypesRepository = &repositories.GetIdentifierTypesRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}
	GetIdentifierTypesUseCase = &usecase.GetIdentifierTypes{Repository: GetIdentifierTypesRepository}
	GetIdentifierTypesClient = &service.GetIdentifierTypesService{UseCase: GetIdentifierTypesUseCase}
}

func buildGetPriceFamilies() {
	if GetPriceFamiliesClient != nil {
		return
	}
	GetPriceFamiliesRepository = &repositories.GetPriceFamiliesRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}
	GetPriceFamiliesUseCase = &usecase.GetPriceFamilies{Repository: GetPriceFamiliesRepository}
	GetPriceFamiliesClient = &service.GetPriceFamiliesService{UseCase: GetPriceFamiliesUseCase}
}

func buildUpdateTransactionByAuthorization() {
	if UpdateTransactionByAuthorizationClient != nil {
		return
	}
	UpdateTransactionByAuthorizationRepository = &repositories.UpdateTransactionByAuthorizationRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}
	UpdateTransactionByAuthorizationUseCase = &usecase.UpdateTransactionByAuthorization{Repository: UpdateTransactionByAuthorizationRepository}
	UpdateTransactionByAuthorizationClient = &service.UpdateTransactionByAuthorizationService{UseCase: UpdateTransactionByAuthorizationUseCase}
}

func ResolveGetDispenserDetailsContainer() iservice.IGetDispenserDetailsService {
	buildGetDispenserDetails()
	return GetDispenserDetailsClient
}

func ResolveGetDispenserDetailsByFamiliesContainer() iservice.IGetDispenserDetailsByFamiliesService {
	buildGetDispenserDetailsByFamilies()
	return GetDispenserDetailsByFamiliesClient
}

func ResolveGetIdentifierTypesContainer() iservice.IGetIdentifierTypesService {
	buildGetIdentifierTypes()
	return GetIdentifierTypesClient
}

func ResolveGetPriceFamiliesContainer() iservice.IGetPriceFamiliesService {
	buildGetPriceFamilies()
	return GetPriceFamiliesClient
}

func ResolveUpdateTransactionByAuthorizationContainer() iservice.IUpdateTransactionByAuthorizationService {
	buildUpdateTransactionByAuthorization()
	return UpdateTransactionByAuthorizationClient
}
