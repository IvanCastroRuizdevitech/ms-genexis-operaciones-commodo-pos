package container_sales

import (
	"ms-genexis-pos-operaciones/context/sales/application/service"
	usecase "ms-genexis-pos-operaciones/context/sales/application/use_case"
	iservice "ms-genexis-pos-operaciones/context/sales/domain/ports/application/service"
	iusecase "ms-genexis-pos-operaciones/context/sales/domain/ports/application/use_case"
	irepositories "ms-genexis-pos-operaciones/context/sales/domain/ports/repositories"
	repositories "ms-genexis-pos-operaciones/context/sales/infrastructure"
	dbclient "ms-genexis-pos-operaciones/infrastructure/db/client"
	presentation_container "ms-genexis-pos-operaciones/presentation/container"
)

var dbConn dbclient.DatabaseConnectionInterface

// REPOSITORIES DB
var CheckPendingSalesRepository irepositories.ICheckPendingSalesRepository
var CheckReadySalesRepository irepositories.ICheckReadySalesRepository
var CheckDatafonoCancellationsInProgressRepository irepositories.ICheckDatafonoCancellationsInProgressRepository
var GetUnresolvedSaleAttributesRepository irepositories.IGetUnresolvedSaleAttributesRepository
var UpdateMovementStateRepository irepositories.IUpdateMovementStateRepository
var AssignCustomerDataRepository irepositories.IAssignCustomerDataRepository
var UpdateClientMovementRepository irepositories.IUpdateClientMovementRepository
var GetPendingSaleDatafonoRepository irepositories.IGetPendingSaleDatafonoRepository
var UpdatePaymentMethodsRepository irepositories.IUpdatePaymentMethodsRepository
var ReprintSaleRepository irepositories.IReprintSaleRepository
var FuelEntryReportRepository irepositories.IFuelEntryReportRepository
var GetDispenserDetailsRepository irepositories.IGetDispenserDetailsRepository

// USECASE
var CheckPendingSalesUseCase iusecase.ICheckPendingSales
var CheckReadySalesUseCase iusecase.ICheckReadySales
var CheckDatafonoCancellationsInProgressUseCase iusecase.ICheckDatafonoCancellationsInProgress
var GetUnresolvedSaleAttributesUseCase iusecase.IGetUnresolvedSaleAttributes
var UpdateMovementStateUseCase iusecase.IUpdateMovementState
var AssignCustomerDataUseCase iusecase.IAssignCustomerData
var UpdateClientMovementUseCase iusecase.IUpdateClientMovement
var GetPendingSaleDatafonoUseCase iusecase.IGetPendingSaleDatafono
var UpdatePaymentMethodsUseCase iusecase.IUpdatePaymentMethods
var ReprintSaleUseCase iusecase.IReprintSale
var FuelEntryReportUseCase iusecase.IFuelEntryReport
var GetDispenserDetailsUseCase iusecase.IGetDispenserDetails

// SERVICE
var CheckPendingSalesClient iservice.ICheckPendingSales
var CheckReadySalesClient iservice.ICheckReadySales
var CheckDatafonoCancellationsInProgressClient iservice.ICheckDatafonoCancellationsInProgress
var GetUnresolvedSaleAttributesClient iservice.IGetUnresolvedSaleAttributes
var UpdateMovementStateClient iservice.IUpdateMovementState
var AssignCustomerDataClient iservice.IAssignCustomerData
var UpdateClientMovementClient iservice.IUpdateClientMovement
var GetPendingSaleDatafonoClient iservice.IGetPendingSaleDatafono
var UpdatePaymentMethodsClient iservice.IUpdatePaymentMethods
var ReprintSaleClient iservice.IReprintSale
var FuelEntryReportClient iservice.IFuelEntryReport
var GetDispenserDetailsClient iservice.IGetDispenserDetails

func resolveDB() dbclient.DatabaseConnectionInterface {
	if dbConn == nil {
		dbConn = presentation_container.ResolveDatabaseConnectionToLecWithPgx()
	}
	return dbConn
}

func buildCheckPendingSales() {
	if CheckPendingSalesClient != nil {
		return
	}
	repo := &repositories.CheckPendingSalesRepository{Connection: resolveDB()}
	CheckPendingSalesRepository = repo
	CheckPendingSalesUseCase = &usecase.CheckPendingSales{Repository: repo}
	CheckPendingSalesClient = &service.CheckPendingSalesClient{UseCase: CheckPendingSalesUseCase}
}

func buildReadySales() {
	if CheckReadySalesClient != nil {
		return
	}
	repo := &repositories.CheckReadySalesRepository{Connection: resolveDB()}
	CheckReadySalesRepository = repo
	CheckReadySalesUseCase = &usecase.CheckReadySales{Repository: repo}
	CheckReadySalesClient = &service.CheckReadySalesClient{UseCase: CheckReadySalesUseCase}
}

func buildDatafonoCancellationsInProgress() {
	if CheckDatafonoCancellationsInProgressClient != nil {
		return
	}
	repo := &repositories.CheckDatafonoCancellationsInProgressRepository{Connection: resolveDB()}
	CheckDatafonoCancellationsInProgressRepository = repo
	CheckDatafonoCancellationsInProgressUseCase = &usecase.CheckDatafonoCancellationsInProgress{Repository: repo}
	CheckDatafonoCancellationsInProgressClient = &service.CheckDatafonoCancellationsInProgressClient{UseCase: CheckDatafonoCancellationsInProgressUseCase}
}

func buildUnresolvedSaleAttributes() {
	if GetUnresolvedSaleAttributesClient != nil {
		return
	}
	repo := &repositories.GetUnresolvedSaleAttributesRepository{Connection: resolveDB()}
	GetUnresolvedSaleAttributesRepository = repo
	GetUnresolvedSaleAttributesUseCase = &usecase.GetUnresolvedSaleAttributes{Repository: repo}
	GetUnresolvedSaleAttributesClient = &service.GetUnresolvedSaleAttributesClient{UseCase: GetUnresolvedSaleAttributesUseCase}
}

func buildUpdateMovementState() {
	if UpdateMovementStateClient != nil {
		return
	}
	repo := &repositories.UpdateMovementStateRepository{Connection: resolveDB()}
	UpdateMovementStateRepository = repo
	UpdateMovementStateUseCase = &usecase.UpdateMovementState{Repository: repo}
	UpdateMovementStateClient = &service.UpdateMovementStateClient{UseCase: UpdateMovementStateUseCase}
}

func buildAssignCustomerData() {
	if AssignCustomerDataClient != nil {
		return
	}
	repo := &repositories.AssignCustomerDataRepository{Connection: resolveDB()}
	AssignCustomerDataRepository = repo
	AssignCustomerDataUseCase = &usecase.AssignCustomerData{Repository: repo}
	AssignCustomerDataClient = &service.AssignCustomerDataClient{UseCase: AssignCustomerDataUseCase}
}

func buildUpdateClientMovement() {
	if UpdateClientMovementClient != nil {
		return
	}
	repo := &repositories.UpdateClientMovementRepository{Connection: resolveDB()}
	UpdateClientMovementRepository = repo
	UpdateClientMovementUseCase = &usecase.UpdateClientMovement{Repository: repo}
	UpdateClientMovementClient = &service.UpdateClientMovementClient{UseCase: UpdateClientMovementUseCase}
}

func buildPendingSaleDatafono() {
	if GetPendingSaleDatafonoClient != nil {
		return
	}
	repo := &repositories.GetPendingSaleDatafonoRepository{Connection: resolveDB()}
	GetPendingSaleDatafonoRepository = repo
	GetPendingSaleDatafonoUseCase = &usecase.GetPendingSaleDatafono{Repository: repo}
	GetPendingSaleDatafonoClient = &service.GetPendingSaleDatafonoClient{UseCase: GetPendingSaleDatafonoUseCase}
}

func buildUpdatePaymentMethods() {
	if UpdatePaymentMethodsClient != nil {
		return
	}
	repo := &repositories.UpdatePaymentMethodsRepository{Connection: resolveDB()}
	UpdatePaymentMethodsRepository = repo
	UpdatePaymentMethodsUseCase = &usecase.UpdatePaymentMethods{Repository: repo}
	UpdatePaymentMethodsClient = &service.UpdatePaymentMethodsClient{UseCase: UpdatePaymentMethodsUseCase}
}

func buildReprintSale() {
	if ReprintSaleClient != nil {
		return
	}
	repo := &repositories.ReprintSaleRepository{Connection: resolveDB()}
	ReprintSaleRepository = repo
	ReprintSaleUseCase = &usecase.ReprintSale{Repository: repo}
	ReprintSaleClient = &service.ReprintSaleClient{UseCase: ReprintSaleUseCase}
}

func buildFuelEntryReport() {
	if FuelEntryReportClient != nil {
		return
	}
	repo := &repositories.FuelEntryReportRepository{Connection: resolveDB()}
	FuelEntryReportRepository = repo
	FuelEntryReportUseCase = &usecase.FuelEntryReport{Repository: repo}
	FuelEntryReportClient = &service.FuelEntryReportClient{UseCase: FuelEntryReportUseCase}
}

func buildGetDispenserDetails() {
	if GetDispenserDetailsClient != nil {
		return
	}
	repo := &repositories.GetDispenserDetailsRepository{Connection: resolveDB()}
	GetDispenserDetailsRepository = repo
	GetDispenserDetailsUseCase = &usecase.GetDispenserDetails{Repository: repo}
	GetDispenserDetailsClient = &service.GetDispenserDetailsClient{UseCase: GetDispenserDetailsUseCase}
}

func ResolveSalesContainer() iservice.ICheckPendingSales {
	buildCheckPendingSales()
	return CheckPendingSalesClient
}

func ResolveReadySalesContainer() iservice.ICheckReadySales {
	buildReadySales()
	return CheckReadySalesClient
}

func ResolveDatafonoCancellationsInProgressContainer() iservice.ICheckDatafonoCancellationsInProgress {
	buildDatafonoCancellationsInProgress()
	return CheckDatafonoCancellationsInProgressClient
}

func ResolveGetUnresolvedSaleAttributesContainer() iservice.IGetUnresolvedSaleAttributes {
	buildUnresolvedSaleAttributes()
	return GetUnresolvedSaleAttributesClient
}

func ResolveUpdateMovementStateContainer() iservice.IUpdateMovementState {
	buildUpdateMovementState()
	return UpdateMovementStateClient
}

func ResolveAssignCustomerDataContainer() iservice.IAssignCustomerData {
	buildAssignCustomerData()
	return AssignCustomerDataClient
}

func ResolveUpdateClientMovementContainer() iservice.IUpdateClientMovement {
	buildUpdateClientMovement()
	return UpdateClientMovementClient
}

func ResolveGetPendingSaleDatafonoContainer() iservice.IGetPendingSaleDatafono {
	buildPendingSaleDatafono()
	return GetPendingSaleDatafonoClient
}

func ResolveUpdatePaymentMethodsContainer() iservice.IUpdatePaymentMethods {
	buildUpdatePaymentMethods()
	return UpdatePaymentMethodsClient
}

func ResolveReprintSaleContainer() iservice.IReprintSale {
	buildReprintSale()
	return ReprintSaleClient
}

func ResolveFuelEntryReportContainer() iservice.IFuelEntryReport {
	buildFuelEntryReport()
	return FuelEntryReportClient
}

func ResolveGetDispenserDetailsContainer() iservice.IGetDispenserDetails {
	buildGetDispenserDetails()
	return GetDispenserDetailsClient
}
