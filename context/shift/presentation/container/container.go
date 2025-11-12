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

// REPOSITORIES HTTPP
var sendOpeningShiftRepositoryHttp irepositories.ISendOpeningShiftRepositoryHttp

// USECASE
var validatePersonShiftUseCase iusecase.IvalidatePersonShift
var openingShiftUseCase iusecase.IOpeningShift

// SERVICE
var openingShift iservice.IOpeningShift

func initializes() {
	getPersonShiftRepository = &repositories.GetPersonShiftRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}

	sendOpeningShiftRepositoryHttp = &repositories.SendOpeningShiftRepositoryHttp{Connection: presentation_container.ResolveClientHttpWithNet()}

	validatePersonShiftUseCase = &usecase.ValidatePersonShift{ShiftRepository: getPersonShiftRepository}
	openingShiftUseCase = &usecase.OpeningShift{SendOpening: sendOpeningShiftRepositoryHttp}

	openingShift = &service.OpeningShiftClient{
		ValidatePerson: validatePersonShiftUseCase,
		OpeningShift:   openingShiftUseCase,
	}

}

func ResolveOpeningShiftContainer() iservice.IOpeningShift {

	if openingShift == nil {
		initializes()
	}
	return openingShift
}
