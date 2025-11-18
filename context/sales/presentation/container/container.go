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
var UpdateMovementStateRepository irepositories.IUpdateMovementStateRepository
var AssignCustomerDataRepository irepositories.IAssignCustomerDataRepository
var UpdateClientMovementRepository irepositories.IUpdateClientMovementRepository
var GetPendingSaleDatafonoRepository irepositories.IGetPendingSaleDatafonoRepository

// USECASE
var CheckPendingSalesUseCase iusecase.ICheckPendingSales
var CheckReadySalesUseCase iusecase.ICheckReadySales
var CheckDatafonoCancellationsInProgressUseCase iusecase.ICheckDatafonoCancellationsInProgress
var GetUnresolvedSaleAttributesUseCase iusecase.IGetUnresolvedSaleAttributes
var UpdateMovementStateUseCase iusecase.IUpdateMovementState
var AssignCustomerDataUseCase iusecase.IAssignCustomerData
var UpdateClientMovementUseCase iusecase.IUpdateClientMovement
var GetPendingSaleDatafonoUseCase iusecase.IGetPendingSaleDatafono

// SERVICE
var CheckPendingSalesClient iservice.ICheckPendingSales
var CheckReadySalesClient iservice.ICheckReadySales
var CheckDatafonoCancellationsInProgressClient iservice.ICheckDatafonoCancellationsInProgress
var GetUnresolvedSaleAttributesClient iservice.IGetUnresolvedSaleAttributes
var UpdateMovementStateClient iservice.IUpdateMovementState
var AssignCustomerDataClient iservice.IAssignCustomerData
var UpdateClientMovementClient iservice.IUpdateClientMovement
var GetPendingSaleDatafonoClient iservice.IGetPendingSaleDatafono

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

    UpdateMovementStateRepository = &repositories.UpdateMovementStateRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}
    UpdateMovementStateUseCase = &usecase.UpdateMovementState{Repository: UpdateMovementStateRepository}
    UpdateMovementStateClient = &service.UpdateMovementStateClient{UseCase: UpdateMovementStateUseCase}

    AssignCustomerDataRepository = &repositories.AssignCustomerDataRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}
    AssignCustomerDataUseCase = &usecase.AssignCustomerData{Repository: AssignCustomerDataRepository}
    AssignCustomerDataClient = &service.AssignCustomerDataClient{UseCase: AssignCustomerDataUseCase}

    UpdateClientMovementRepository = &repositories.UpdateClientMovementRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}
    UpdateClientMovementUseCase = &usecase.UpdateClientMovement{Repository: UpdateClientMovementRepository}
    UpdateClientMovementClient = &service.UpdateClientMovementClient{UseCase: UpdateClientMovementUseCase}

    GetPendingSaleDatafonoRepository = &repositories.GetPendingSaleDatafonoRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}
    GetPendingSaleDatafonoUseCase = &usecase.GetPendingSaleDatafono{Repository: GetPendingSaleDatafonoRepository}
    GetPendingSaleDatafonoClient = &service.GetPendingSaleDatafonoClient{UseCase: GetPendingSaleDatafonoUseCase}
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

func ResolveUpdateMovementStateContainer() iservice.IUpdateMovementState {
    if UpdateMovementStateClient == nil {
        initializes()
    }
    return UpdateMovementStateClient
}

func ResolveAssignCustomerDataContainer() iservice.IAssignCustomerData {
    if AssignCustomerDataClient == nil {
        initializes()
    }
    return AssignCustomerDataClient
}

func ResolveUpdateClientMovementContainer() iservice.IUpdateClientMovement {
    if UpdateClientMovementClient == nil {
        initializes()
    }
    return UpdateClientMovementClient
}

func ResolveGetPendingSaleDatafonoContainer() iservice.IGetPendingSaleDatafono {
    if GetPendingSaleDatafonoClient == nil {
        initializes()
    }
    return GetPendingSaleDatafonoClient
}
