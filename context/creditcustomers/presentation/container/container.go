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

var CreditCustomerRepository irepositories.ICreditCustomerRepository

// USE CASES
var GetCreditCustomersUseCase iusecase.IGetCreditCustomers
var UpdateCreditCustomerLimitUseCase iusecase.IUpdateCreditCustomerLimit

// SERVICES
var GetCreditCustomersClient iservice.IGetCreditCustomersService
var UpdateCreditCustomerLimitClient iservice.IUpdateCreditCustomerLimitService

func resolveDB() dbclient.DatabaseConnectionInterface {
	if dbConn == nil {
		dbConn = presentation_container.ResolveDatabaseConnectionToLecWithPgx()
	}
	return dbConn
}

func ensureRepository() irepositories.ICreditCustomerRepository {
	if CreditCustomerRepository == nil {
		CreditCustomerRepository = &repositories.CreditCustomerRepository{
			Connection: resolveDB(),
		}
	}
	return CreditCustomerRepository
}

func buildGetCreditCustomers() {
	if GetCreditCustomersClient != nil {
		return
	}
	repo := ensureRepository()
	GetCreditCustomersUseCase = &usecase.GetCreditCustomers{Repository: repo}
	GetCreditCustomersClient = &app_service.GetCreditCustomersService{UseCase: GetCreditCustomersUseCase}
}

func buildUpdateCreditCustomerLimit() {
	if UpdateCreditCustomerLimitClient != nil {
		return
	}
	repo := ensureRepository()
	UpdateCreditCustomerLimitUseCase = &usecase.UpdateCreditCustomerLimit{Repository: repo}
	UpdateCreditCustomerLimitClient = &app_service.UpdateCreditCustomerLimitService{UseCase: UpdateCreditCustomerLimitUseCase}
}

func ResolveGetCreditCustomersContainer() iservice.IGetCreditCustomersService {
	buildGetCreditCustomers()
	return GetCreditCustomersClient
}

func ResolveUpdateCreditCustomerLimitContainer() iservice.IUpdateCreditCustomerLimitService {
	buildUpdateCreditCustomerLimit()
	return UpdateCreditCustomerLimitClient
}
