package container_sales

import (
    "ms-genexis-pos-operaciones/context/sales/application/service"
    usecase "ms-genexis-pos-operaciones/context/sales/application/use_case"
    iservice "ms-genexis-pos-operaciones/context/sales/domain/ports/application/service"
    iusecase "ms-genexis-pos-operaciones/context/sales/domain/ports/application/use_case"
    irepositories "ms-genexis-pos-operaciones/context/sales/domain/ports/repositories"
    repositories "ms-genexis-pos-operaciones/context/sales/infrastructure"
    presentation_container "ms-genexis-pos-operaciones/presentation/container"
)

// REPOSITORIES DB
var CheckPendingSalesRepository irepositories.ICheckPendingSalesRepository
var CheckReadySalesRepository irepositories.ICheckReadySalesRepository
var CheckDatafonoCancellationsInProgressRepository irepositories.ICheckDatafonoCancellationsInProgressRepository
var GetUnresolvedSaleAttributesRepository irepositories.IGetUnresolvedSaleAttributesRepository

// USECASE
var CheckPendingSalesUseCase iusecase.ICheckPendingSales
var CheckReadySalesUseCase iusecase.ICheckReadySales
var CheckDatafonoCancellationsInProgressUseCase iusecase.ICheckDatafonoCancellationsInProgress
var GetUnresolvedSaleAttributesUseCase iusecase.IGetUnresolvedSaleAttributes

// SERVICE
var CheckPendingSalesClient iservice.ICheckPendingSales
var CheckReadySalesClient iservice.ICheckReadySales
var CheckDatafonoCancellationsInProgressClient iservice.ICheckDatafonoCancellationsInProgress
var GetUnresolvedSaleAttributesClient iservice.IGetUnresolvedSaleAttributes

func initializes() {
    CheckPendingSalesRepository = &repositories.CheckPendingSalesRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}
    CheckPendingSalesUseCase = &usecase.CheckPendingSales{Repository: CheckPendingSalesRepository}
    CheckPendingSalesClient = &service.CheckPendingSalesClient{UseCase: CheckPendingSalesUseCase}

    CheckReadySalesRepository = &repositories.CheckReadySalesRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}
    CheckReadySalesUseCase = &usecase.CheckReadySales{Repository: CheckReadySalesRepository}
    CheckReadySalesClient = &service.CheckReadySalesClient{UseCase: CheckReadySalesUseCase}

    CheckDatafonoCancellationsInProgressRepository = &repositories.CheckDatafonoCancellationsInProgressRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}
    CheckDatafonoCancellationsInProgressUseCase = &usecase.CheckDatafonoCancellationsInProgress{Repository: CheckDatafonoCancellationsInProgressRepository}
    CheckDatafonoCancellationsInProgressClient = &service.CheckDatafonoCancellationsInProgressClient{UseCase: CheckDatafonoCancellationsInProgressUseCase}

    GetUnresolvedSaleAttributesRepository = &repositories.GetUnresolvedSaleAttributesRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}
    GetUnresolvedSaleAttributesUseCase = &usecase.GetUnresolvedSaleAttributes{Repository: GetUnresolvedSaleAttributesRepository}
    GetUnresolvedSaleAttributesClient = &service.GetUnresolvedSaleAttributesClient{UseCase: GetUnresolvedSaleAttributesUseCase}
}

func ResolveSalesContainer() iservice.ICheckPendingSales {
    if CheckPendingSalesClient == nil {
        initializes()
    }
    return CheckPendingSalesClient
}

func ResolveReadySalesContainer() iservice.ICheckReadySales {
    if CheckReadySalesClient == nil {
        initializes()
    }
    return CheckReadySalesClient
}

func ResolveDatafonoCancellationsInProgressContainer() iservice.ICheckDatafonoCancellationsInProgress {
    if CheckDatafonoCancellationsInProgressClient == nil {
        initializes()
    }
    return CheckDatafonoCancellationsInProgressClient
}

func ResolveGetUnresolvedSaleAttributesContainer() iservice.IGetUnresolvedSaleAttributes {
    if GetUnresolvedSaleAttributesClient == nil {
        initializes()
    }
    return GetUnresolvedSaleAttributesClient
}

