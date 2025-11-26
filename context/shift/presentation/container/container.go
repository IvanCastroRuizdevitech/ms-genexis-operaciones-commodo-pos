package container_shift

import (
	"ms-genexis-pos-operaciones/context/shift/application/service"
	usecase "ms-genexis-pos-operaciones/context/shift/application/use_case"
	iservice "ms-genexis-pos-operaciones/context/shift/domain/ports/application/service"
	iusecase "ms-genexis-pos-operaciones/context/shift/domain/ports/application/use_case"
	irepositories "ms-genexis-pos-operaciones/context/shift/domain/ports/repositories"
	repositories "ms-genexis-pos-operaciones/context/shift/infrastructure"
	dbclient "ms-genexis-pos-operaciones/infrastructure/db/client"
	externalhttp "ms-genexis-pos-operaciones/infrastructure/externals/externalhttp"
	presentation_container "ms-genexis-pos-operaciones/presentation/container"
)

// REPOSITORIES DB
var getPersonShiftRepository irepositories.IGetPersonShiftRepository
var getDailyIncomeMeasurementsRepository irepositories.IGetDailyIncomeMeasurementsRepository
var getFuelPumpsRepository irepositories.IGetFuelPumpsRepository
var personValidationRepository irepositories.IValidatePersonRepository

// REPOSITORIES HTTPP
var sendOpeningShiftRepositoryHttp irepositories.ISendOpeningShiftRepositoryHttp

// USECASE
var validatePersonShiftUseCase iusecase.IvalidatePersonShift
var openingShiftUseCase iusecase.IOpeningShift
var getDailyIncomeMeasurementsUseCase iusecase.IGetDailyIncomeMeasurements
var getFuelPumpsUseCase iusecase.IGetFuelPumps
var personValidationUseCase iusecase.IPersonValidation

// SERVICE
var openingShift iservice.IOpeningShift
var dailyIncomeMeasurementsService *service.DailyIncomeMeasurementsClient
var fuelPumpsService *service.FuelPumpsClient
var personValidationService *service.PersonValidationClient

func resolveDB() dbclient.DatabaseConnectionInterface {
	return presentation_container.ResolveDatabaseConnectionToLecWithPgx()
}

func resolveHTTPClient() externalhttp.ClientHTTPInterface {
	return presentation_container.ResolveClientHttpWithNet()
}

func ensurePersonRepositories(dbConn dbclient.DatabaseConnectionInterface) {
	if getPersonShiftRepository != nil {
		return
	}
	repo := &repositories.GetPersonShiftRepository{Connection: dbConn}
	getPersonShiftRepository = repo
	personValidationRepository = repo
}

func ensureDailyIncomeMeasurementsRepository(dbConn dbclient.DatabaseConnectionInterface) {
	if getDailyIncomeMeasurementsRepository == nil {
		getDailyIncomeMeasurementsRepository = &repositories.GetDailyIncomeMeasurementsRepository{Connection: dbConn}
	}
}

func ensureFuelPumpsRepository(dbConn dbclient.DatabaseConnectionInterface) {
	if getFuelPumpsRepository == nil {
		getFuelPumpsRepository = &repositories.GetFuelPumpsRepository{Connection: dbConn}
	}
}

func buildOpeningShift() {
	if openingShift != nil {
		return
	}
	dbConn := resolveDB()
	ensurePersonRepositories(dbConn)
	if validatePersonShiftUseCase == nil {
		validatePersonShiftUseCase = &usecase.ValidatePersonShift{ShiftRepository: getPersonShiftRepository}
	}
	if sendOpeningShiftRepositoryHttp == nil {
		sendOpeningShiftRepositoryHttp = &repositories.SendOpeningShiftRepositoryHttp{Connection: resolveHTTPClient()}
	}
	if openingShiftUseCase == nil {
		openingShiftUseCase = &usecase.OpeningShift{SendOpening: sendOpeningShiftRepositoryHttp}
	}
	openingShift = &service.OpeningShiftClient{
		ValidatePerson: validatePersonShiftUseCase,
		OpeningShift:   openingShiftUseCase,
	}
}

func buildDailyIncomeMeasurements() {
	if dailyIncomeMeasurementsService != nil {
		return
	}
	dbConn := resolveDB()
	ensureDailyIncomeMeasurementsRepository(dbConn)
	if getDailyIncomeMeasurementsUseCase == nil {
		getDailyIncomeMeasurementsUseCase = &usecase.GetDailyIncomeMeasurements{Repository: getDailyIncomeMeasurementsRepository}
	}
	dailyIncomeMeasurementsService = &service.DailyIncomeMeasurementsClient{
		GetDailyIncomeMeasurements: getDailyIncomeMeasurementsUseCase,
	}
}

func buildFuelPumps() {
	if fuelPumpsService != nil {
		return
	}
	dbConn := resolveDB()
	ensureFuelPumpsRepository(dbConn)
	if getFuelPumpsUseCase == nil {
		getFuelPumpsUseCase = &usecase.GetFuelPumps{Repository: getFuelPumpsRepository}
	}
	fuelPumpsService = &service.FuelPumpsClient{GetFuelPumps: getFuelPumpsUseCase}
}

func buildPersonValidation() {
	if personValidationService != nil {
		return
	}
	dbConn := resolveDB()
	ensurePersonRepositories(dbConn)
	if personValidationUseCase == nil {
		personValidationUseCase = &usecase.PersonValidation{Repository: personValidationRepository}
	}
	personValidationService = &service.PersonValidationClient{ValidatePerson: personValidationUseCase}
}

func ResolveOpeningShiftContainer() iservice.IOpeningShift {
	buildOpeningShift()
	return openingShift
}

func ResolveDailyIncomeMeasurementsContainer() *service.DailyIncomeMeasurementsClient {
	buildDailyIncomeMeasurements()
	return dailyIncomeMeasurementsService
}

func ResolveFuelPumpsContainer() *service.FuelPumpsClient {
	buildFuelPumps()
	return fuelPumpsService
}

func ResolvePersonValidationContainer() *service.PersonValidationClient {
	buildPersonValidation()
	return personValidationService
}
