package container_shift

import (
    "ms-genexis-pos-operaciones/context/shift/application/service"
    usecase "ms-genexis-pos-operaciones/context/shift/application/use_case"
    iservice "ms-genexis-pos-operaciones/context/shift/domain/ports/application/service"
    iusecase "ms-genexis-pos-operaciones/context/shift/domain/ports/application/use_case"
    irepositories "ms-genexis-pos-operaciones/context/shift/domain/ports/repositories"
    repositories "ms-genexis-pos-operaciones/context/shift/infrastructure"
    presentation_container "ms-genexis-pos-operaciones/presentation/container"
)

// REPOSITORIES DB
var getPersonShiftRepository irepositories.IGetPersonShiftRepository
var getDailyIncomeMeasurementsRepository irepositories.IGetDailyIncomeMeasurementsRepository

// REPOSITORIES HTTPP
var sendOpeningShiftRepositoryHttp irepositories.ISendOpeningShiftRepositoryHttp

// USECASE
var validatePersonShiftUseCase iusecase.IvalidatePersonShift
var openingShiftUseCase iusecase.IOpeningShift
var getDailyIncomeMeasurementsUseCase iusecase.IGetDailyIncomeMeasurements

// SERVICE
var openingShift iservice.IOpeningShift
var dailyIncomeMeasurementsService *service.DailyIncomeMeasurementsClient

func initializes() {
    getPersonShiftRepository = &repositories.GetPersonShiftRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}
    getDailyIncomeMeasurementsRepository = &repositories.GetDailyIncomeMeasurementsRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}

	sendOpeningShiftRepositoryHttp = &repositories.SendOpeningShiftRepositoryHttp{Connection: presentation_container.ResolveClientHttpWithNet()}

    validatePersonShiftUseCase = &usecase.ValidatePersonShift{ShiftRepository: getPersonShiftRepository}
    openingShiftUseCase = &usecase.OpeningShift{SendOpening: sendOpeningShiftRepositoryHttp}
    getDailyIncomeMeasurementsUseCase = &usecase.GetDailyIncomeMeasurements{Repository: getDailyIncomeMeasurementsRepository}

    openingShift = &service.OpeningShiftClient{
        ValidatePerson: validatePersonShiftUseCase,
        OpeningShift:   openingShiftUseCase,
    }

    dailyIncomeMeasurementsService = &service.DailyIncomeMeasurementsClient{
        GetDailyIncomeMeasurements: getDailyIncomeMeasurementsUseCase,
    }

}

func ResolveOpeningShiftContainer() iservice.IOpeningShift {

	if openingShift == nil {
		initializes()
	}
    return openingShift
}

func ResolveDailyIncomeMeasurementsContainer() *service.DailyIncomeMeasurementsClient {
    if dailyIncomeMeasurementsService == nil {
        initializes()
    }
    return dailyIncomeMeasurementsService
}
