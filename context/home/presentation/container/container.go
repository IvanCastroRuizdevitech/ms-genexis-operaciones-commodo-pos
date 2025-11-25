package container_home

import (
	"ms-genexis-pos-operaciones/context/home/application/service"
	usecase "ms-genexis-pos-operaciones/context/home/application/use_case"
	iservice "ms-genexis-pos-operaciones/context/home/domain/ports/application/service"
	iusecase "ms-genexis-pos-operaciones/context/home/domain/ports/application/use_case"
	irepositories "ms-genexis-pos-operaciones/context/home/domain/ports/repositories"
	repositories "ms-genexis-pos-operaciones/context/home/infrastructure"
	dbclient "ms-genexis-pos-operaciones/infrastructure/db/client"
	presentation_container "ms-genexis-pos-operaciones/presentation/container"
)

var loadErrorNotificationRepository irepositories.ILoadErrorNotificationRepository
var loadErrorNotificationUseCase iusecase.ILoadErrorNotification
var loadErrorNotificationService iservice.ILoadErrorNotification

var pendingTransmissionsRepository irepositories.IGetPendingTransmissionsRepository
var pendingTransmissionsUseCase iusecase.IGetPendingTransmissions
var pendingTransmissionsService iservice.IPendingTransmissions

var updateTransmissionStatusRepository irepositories.IUpdateTransmissionStatusRepository
var processPendingTransmissionsUseCase iusecase.IProcessPendingTransmissions
var processPendingTransmissionsService iservice.IPendingTransmissionsProcessor

func resolveDB() dbclient.DatabaseConnectionInterface {
	return presentation_container.ResolveDatabaseConnectionToLecWithPgx()
}

func buildLoadErrorNotification() {
	if loadErrorNotificationService != nil {
		return
	}
	dbConn := resolveDB()
	loadErrorNotificationRepository = &repositories.LoadErrorNotificationRepository{Connection: dbConn}
	loadErrorNotificationUseCase = &usecase.LoadErrorNotification{Repository: loadErrorNotificationRepository}
	loadErrorNotificationService = &service.LoadErrorNotificationClient{LoadErrorNotification: loadErrorNotificationUseCase}
}

func buildPendingTransmissions() {
	if pendingTransmissionsService != nil {
		return
	}
	dbConn := resolveDB()
	pendingTransmissionsRepository = &repositories.GetPendingTransmissionsRepository{Connection: dbConn}
	pendingTransmissionsUseCase = &usecase.GetPendingTransmissions{Repository: pendingTransmissionsRepository}
	pendingTransmissionsService = &service.PendingTransmissionsClient{GetPendingTransmissions: pendingTransmissionsUseCase}
}

func buildProcessPendingTransmissions() {
	if processPendingTransmissionsService != nil {
		return
	}
	dbConn := resolveDB()
	if pendingTransmissionsRepository == nil {
		pendingTransmissionsRepository = &repositories.GetPendingTransmissionsRepository{Connection: dbConn}
		pendingTransmissionsUseCase = &usecase.GetPendingTransmissions{Repository: pendingTransmissionsRepository}
		pendingTransmissionsService = &service.PendingTransmissionsClient{GetPendingTransmissions: pendingTransmissionsUseCase}
	}
	updateTransmissionStatusRepository = &repositories.UpdateTransmissionStatusRepository{Connection: dbConn}
	processPendingTransmissionsUseCase = &usecase.ProcessPendingTransmissions{
		FetchRepository:  pendingTransmissionsRepository,
		UpdateRepository: updateTransmissionStatusRepository,
		HTTPClient:       presentation_container.ResolveClientHttpWithNet(),
	}
	processPendingTransmissionsService = &service.PendingTransmissionsProcessor{ProcessPendingTransmissions: processPendingTransmissionsUseCase}
}

func ResolveLoadErrorNotificationContainer() iservice.ILoadErrorNotification {
	buildLoadErrorNotification()
	return loadErrorNotificationService
}

func ResolvePendingTransmissionsContainer() iservice.IPendingTransmissions {
	buildPendingTransmissions()
	return pendingTransmissionsService
}

func ResolveProcessPendingTransmissionsContainer() iservice.IPendingTransmissionsProcessor {
	buildProcessPendingTransmissions()
	return processPendingTransmissionsService
}
