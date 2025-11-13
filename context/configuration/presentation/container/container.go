package container_configuration

import (
    "ms-genexis-pos-operaciones/context/configuration/application/service"
    usecase "ms-genexis-pos-operaciones/context/configuration/application/use_case"
    iservice "ms-genexis-pos-operaciones/context/configuration/domain/ports/application/service"
    iusecase "ms-genexis-pos-operaciones/context/configuration/domain/ports/application/use_case"
    irepositories "ms-genexis-pos-operaciones/context/configuration/domain/ports/repositories"
    repositories "ms-genexis-pos-operaciones/context/configuration/infrastructure"
    presentation_container "ms-genexis-pos-operaciones/presentation/container"
)

// REPOSITORIES DB
var GetParametersRepository irepositories.IGetParametersRepository
var GetPromoterDutyRepository irepositories.IGetPromoterDutyRepository

// REPOSITORIES HTTPP

// USECASE
var GetParametersUseCase iusecase.IGetParameters
var GetPromoterDutyUseCase iusecase.IGetPromoterDuty

// SERVICE
var GetParametersClient iservice.IGetParameters
var GetPromoterDutyClient iservice.IGetPromoterDuty

func initializes() {
    GetParametersRepository = &repositories.GetParametersRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}
    GetParametersUseCase = &usecase.GetParameters{GetParameters: GetParametersRepository}
    GetParametersClient = &service.GetParametersClient{GetParameters: GetParametersUseCase}

    GetPromoterDutyRepository = &repositories.GetPromoterDutyRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}
    GetPromoterDutyUseCase = &usecase.GetPromoterDuty{GetPromoterDuty: GetPromoterDutyRepository}
    GetPromoterDutyClient = &service.GetPromoterDutyClient{GetPromoterDuty: GetPromoterDutyUseCase}
}

func ResolveGetParametersContainer() iservice.IGetParameters {

	if GetParametersClient == nil {
		initializes()
	}

    return GetParametersClient
}

func ResolveGetPromoterDutyContainer() iservice.IGetPromoterDuty {
    if GetPromoterDutyClient == nil {
        initializes()
    }
    return GetPromoterDutyClient
}
