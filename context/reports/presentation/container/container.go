package container_reports

import (
	"ms-genexis-pos-operaciones/context/reports/application/service"
	usecase "ms-genexis-pos-operaciones/context/reports/application/use_case"
	iservice "ms-genexis-pos-operaciones/context/reports/domain/ports/application/service"
	iusecase "ms-genexis-pos-operaciones/context/reports/domain/ports/application/use_case"
	irepositories "ms-genexis-pos-operaciones/context/reports/domain/ports/repositories"
	repositories "ms-genexis-pos-operaciones/context/reports/infrastructure"
	dbclient "ms-genexis-pos-operaciones/infrastructure/db/client"
	presentation_container "ms-genexis-pos-operaciones/presentation/container"
)

// REPOSITORIES DB
var GetDayClosingReportRepository irepositories.IGetDayClosingReportRepository
var GetFuelReportRepository irepositories.IGetFuelReportRepository
var GetDailyNoveltiesRepository irepositories.IGetDailyNoveltiesRepository
var CreateTankPrintEventRepository irepositories.ICreateTankPrintEventRepository
var GetMovementTypesRepository irepositories.IGetMovementTypesRepository
var GetTanksRepository irepositories.IGetTanksRepository

// USECASE
var GetDayClosingReportUseCase iusecase.IGetDayClosingReport
var GetFuelReportUseCase iusecase.IGetFuelReport
var GetDailyNoveltiesUseCase iusecase.IGetDailyNovelties
var CreateTankPrintEventUseCase iusecase.ICreateTankPrintEvent
var GetMovementTypesUseCase iusecase.IGetMovementTypes
var GetTanksUseCase iusecase.IGetTanks

// SERVICE
var DayClosingReportClient iservice.IDayClosingReport
var FuelReportClient iservice.IFuelReport
var DailyNoveltiesClient iservice.IDailyNovelties
var TankPrintEventClient iservice.ITankPrintEvent
var MovementTypesClient iservice.IMovementTypes
var TanksClient iservice.ITanks

func resolveDB() dbclient.DatabaseConnectionInterface {
	return presentation_container.ResolveDatabaseConnectionToLecWithPgx()
}

func buildDayClosingReport() {
	if DayClosingReportClient != nil {
		return
	}
	dbConn := resolveDB()
	GetDayClosingReportRepository = &repositories.GetDayClosingReportRepository{Connection: dbConn}
	GetDayClosingReportUseCase = &usecase.GetDayClosingReport{Repository: GetDayClosingReportRepository}
	DayClosingReportClient = &service.DayClosingReportClient{GetDayClosingReport: GetDayClosingReportUseCase}
}

func buildFuelReport() {
	if FuelReportClient != nil {
		return
	}
	dbConn := resolveDB()
	GetFuelReportRepository = &repositories.GetFuelReportRepository{Connection: dbConn}
	GetFuelReportUseCase = &usecase.GetFuelReport{Repository: GetFuelReportRepository}
	FuelReportClient = &service.FuelReportClient{GetFuelReport: GetFuelReportUseCase}
}

func buildDailyNovelties() {
	if DailyNoveltiesClient != nil {
		return
	}
	dbConn := resolveDB()
	GetDailyNoveltiesRepository = &repositories.GetDailyNoveltiesRepository{Connection: dbConn}
	GetDailyNoveltiesUseCase = &usecase.GetDailyNovelties{Repository: GetDailyNoveltiesRepository}
	DailyNoveltiesClient = &service.DailyNoveltiesClient{GetDailyNovelties: GetDailyNoveltiesUseCase}
}

func buildTankPrintEvent() {
	if TankPrintEventClient != nil {
		return
	}
	dbConn := resolveDB()
	CreateTankPrintEventRepository = &repositories.CreateTankPrintEventRepository{Connection: dbConn}
	CreateTankPrintEventUseCase = &usecase.CreateTankPrintEvent{Repository: CreateTankPrintEventRepository}
	TankPrintEventClient = &service.TankPrintEventClient{CreateTankPrintEvent: CreateTankPrintEventUseCase}
}

func buildMovementTypes() {
	if MovementTypesClient != nil {
		return
	}
	dbConn := resolveDB()
	GetMovementTypesRepository = &repositories.GetMovementTypesRepository{Connection: dbConn}
	GetMovementTypesUseCase = &usecase.GetMovementTypes{Repository: GetMovementTypesRepository}
	MovementTypesClient = &service.MovementTypesClient{GetMovementTypes: GetMovementTypesUseCase}
}

func buildTanks() {
	if TanksClient != nil {
		return
	}
	dbConn := resolveDB()
	GetTanksRepository = &repositories.GetTanksRepository{Connection: dbConn}
	GetTanksUseCase = &usecase.GetTanks{Repository: GetTanksRepository}
	TanksClient = &service.TanksClient{GetTanks: GetTanksUseCase}
}

func ResolveDayClosingReportContainer() iservice.IDayClosingReport {
	buildDayClosingReport()
	return DayClosingReportClient
}

func ResolveFuelReportContainer() iservice.IFuelReport {
	buildFuelReport()
	return FuelReportClient
}

func ResolveDailyNoveltiesContainer() iservice.IDailyNovelties {
	buildDailyNovelties()
	return DailyNoveltiesClient
}

func ResolveTankPrintEventContainer() iservice.ITankPrintEvent {
	buildTankPrintEvent()
	return TankPrintEventClient
}

func ResolveMovementTypesContainer() iservice.IMovementTypes {
	buildMovementTypes()
	return MovementTypesClient
}

func ResolveTanksContainer() iservice.ITanks {
	buildTanks()
	return TanksClient
}
