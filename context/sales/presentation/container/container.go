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

// USECASE
var CheckPendingSalesUseCase iusecase.ICheckPendingSales

// SERVICE
var CheckPendingSalesClient iservice.ICheckPendingSales

func initializes() {
    CheckPendingSalesRepository = &repositories.CheckPendingSalesRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}
    CheckPendingSalesUseCase = &usecase.CheckPendingSales{Repository: CheckPendingSalesRepository}
    CheckPendingSalesClient = &service.CheckPendingSalesClient{UseCase: CheckPendingSalesUseCase}
}

func ResolveSalesContainer() iservice.ICheckPendingSales {
    if CheckPendingSalesClient == nil {
        initializes()
    }
    return CheckPendingSalesClient
}

