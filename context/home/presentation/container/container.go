package container_home

import (
    "ms-genexis-pos-operaciones/context/home/application/service"
    usecase "ms-genexis-pos-operaciones/context/home/application/use_case"
    iservice "ms-genexis-pos-operaciones/context/home/domain/ports/application/service"
    iusecase "ms-genexis-pos-operaciones/context/home/domain/ports/application/use_case"
    irepositories "ms-genexis-pos-operaciones/context/home/domain/ports/repositories"
    repositories "ms-genexis-pos-operaciones/context/home/infrastructure"
    presentation_container "ms-genexis-pos-operaciones/presentation/container"
)

var loadErrorNotificationRepository irepositories.ILoadErrorNotificationRepository
var loadErrorNotificationUseCase iusecase.ILoadErrorNotification
var loadErrorNotificationService iservice.ILoadErrorNotification

func initializes() {
    loadErrorNotificationRepository = &repositories.LoadErrorNotificationRepository{
        Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx(),
    }
    loadErrorNotificationUseCase = &usecase.LoadErrorNotification{
        Repository: loadErrorNotificationRepository,
    }
    loadErrorNotificationService = &service.LoadErrorNotificationClient{
        LoadErrorNotification: loadErrorNotificationUseCase,
    }
}

func ResolveLoadErrorNotificationContainer() iservice.ILoadErrorNotification {
    if loadErrorNotificationService == nil {
        initializes()
    }
    return loadErrorNotificationService
}
