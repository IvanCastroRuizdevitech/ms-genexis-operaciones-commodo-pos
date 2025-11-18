package container_reports

import (
	"ms-genexis-pos-operaciones/context/reports/application/service"
	usecase "ms-genexis-pos-operaciones/context/reports/application/use_case"
	iservice "ms-genexis-pos-operaciones/context/reports/domain/ports/application/service"
	iusecase "ms-genexis-pos-operaciones/context/reports/domain/ports/application/use_case"
	irepositories "ms-genexis-pos-operaciones/context/reports/domain/ports/repositories"
	repositories "ms-genexis-pos-operaciones/context/reports/infrastructure"
	presentation_container "ms-genexis-pos-operaciones/presentation/container"
)

// REPOSITORIES DB
var GetDayClosingReportRepository irepositories.IGetDayClosingReportRepository
var GetFuelReportRepository irepositories.IGetFuelReportRepository
var GetDailyNoveltiesRepository irepositories.IGetDailyNoveltiesRepository
var CreateTankPrintEventRepository irepositories.ICreateTankPrintEventRepository

// USECASE
var GetDayClosingReportUseCase iusecase.IGetDayClosingReport
var GetFuelReportUseCase iusecase.IGetFuelReport
var GetDailyNoveltiesUseCase iusecase.IGetDailyNovelties
var CreateTankPrintEventUseCase iusecase.ICreateTankPrintEvent

// SERVICE
var DayClosingReportClient iservice.IDayClosingReport
var FuelReportClient iservice.IFuelReport
var DailyNoveltiesClient iservice.IDailyNovelties
var TankPrintEventClient iservice.ITankPrintEvent

func initializes() {
	GetDayClosingReportRepository = &repositories.GetDayClosingReportRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}
	GetFuelReportRepository = &repositories.GetFuelReportRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}
	GetDailyNoveltiesRepository = &repositories.GetDailyNoveltiesRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}
	CreateTankPrintEventRepository = &repositories.CreateTankPrintEventRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}

	GetDayClosingReportUseCase = &usecase.GetDayClosingReport{Repository: GetDayClosingReportRepository}
	GetFuelReportUseCase = &usecase.GetFuelReport{Repository: GetFuelReportRepository}
	GetDailyNoveltiesUseCase = &usecase.GetDailyNovelties{Repository: GetDailyNoveltiesRepository}
	CreateTankPrintEventUseCase = &usecase.CreateTankPrintEvent{Repository: CreateTankPrintEventRepository}

	DayClosingReportClient = &service.DayClosingReportClient{GetDayClosingReport: GetDayClosingReportUseCase}
	FuelReportClient = &service.FuelReportClient{GetFuelReport: GetFuelReportUseCase}
	DailyNoveltiesClient = &service.DailyNoveltiesClient{GetDailyNovelties: GetDailyNoveltiesUseCase}
	TankPrintEventClient = &service.TankPrintEventClient{CreateTankPrintEvent: CreateTankPrintEventUseCase}
}

func ResolveDayClosingReportContainer() iservice.IDayClosingReport {
	if DayClosingReportClient == nil {
		initializes()
	}
	return DayClosingReportClient
}

func ResolveFuelReportContainer() iservice.IFuelReport {
	if FuelReportClient == nil {
		initializes()
	}
	return FuelReportClient
}

func ResolveDailyNoveltiesContainer() iservice.IDailyNovelties {
	if DailyNoveltiesClient == nil {
		initializes()
	}
	return DailyNoveltiesClient
}

func ResolveTankPrintEventContainer() iservice.ITankPrintEvent {
	if TankPrintEventClient == nil {
		initializes()
	}
	return TankPrintEventClient
}
