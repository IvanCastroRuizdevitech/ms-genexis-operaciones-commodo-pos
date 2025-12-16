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

func ResolveUpdateTransactionByAuthorizationContainer() iservice.IUpdateTransactionByAuthorizationService {
	buildUpdateTransactionByAuthorization()
	return UpdateTransactionByAuthorizationClient
}
