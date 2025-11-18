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

// USECASE
var GetDayClosingReportUseCase iusecase.IGetDayClosingReport
var GetFuelReportUseCase iusecase.IGetFuelReport

// SERVICE
var DayClosingReportClient iservice.IDayClosingReport
var FuelReportClient iservice.IFuelReport

func initializes() {
	GetDayClosingReportRepository = &repositories.GetDayClosingReportRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}
	GetFuelReportRepository = &repositories.GetFuelReportRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}

	GetDayClosingReportUseCase = &usecase.GetDayClosingReport{Repository: GetDayClosingReportRepository}
	GetFuelReportUseCase = &usecase.GetFuelReport{Repository: GetFuelReportRepository}

	DayClosingReportClient = &service.DayClosingReportClient{GetDayClosingReport: GetDayClosingReportUseCase}
	FuelReportClient = &service.FuelReportClient{GetFuelReport: GetFuelReportUseCase}
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
