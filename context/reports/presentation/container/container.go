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

// USECASE
var GetDayClosingReportUseCase iusecase.IGetDayClosingReport

// SERVICE
var DayClosingReportClient iservice.IDayClosingReport

func initializes() {
    GetDayClosingReportRepository = &repositories.GetDayClosingReportRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}
    GetDayClosingReportUseCase = &usecase.GetDayClosingReport{Repository: GetDayClosingReportRepository}
    DayClosingReportClient = &service.DayClosingReportClient{GetDayClosingReport: GetDayClosingReportUseCase}
}

func ResolveDayClosingReportContainer() iservice.IDayClosingReport {
    if DayClosingReportClient == nil {
        initializes()
    }
    return DayClosingReportClient
}

