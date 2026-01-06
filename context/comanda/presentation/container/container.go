package container_comanda

import (
	"ms-genexis-pos-operaciones/context/comanda/application/service"
	usecase "ms-genexis-pos-operaciones/context/comanda/application/use_case"
	iservice "ms-genexis-pos-operaciones/context/comanda/domain/ports/application/service"
	iusecase "ms-genexis-pos-operaciones/context/comanda/domain/ports/application/use_case"
	irepositories "ms-genexis-pos-operaciones/context/comanda/domain/ports/repositories"
	repositories "ms-genexis-pos-operaciones/context/comanda/infrastructure"
	dbclient "ms-genexis-pos-operaciones/infrastructure/db/client"
	presentation_container "ms-genexis-pos-operaciones/presentation/container"
)

var dbConn dbclient.DatabaseConnectionInterface

// REPOSITORIES DB
var UpdateComandaStatusRepository irepositories.IUpdateComandaStatusRepository

// USECASE
var UpdateComandaStatusUseCase iusecase.IUpdateComandaStatus

// SERVICE
var UpdateComandaStatusClient iservice.IUpdateComandaStatus

func resolveDB() dbclient.DatabaseConnectionInterface {
	if dbConn == nil {
		dbConn = presentation_container.ResolveDatabaseConnectionToLecWithPgx()
	}
	return dbConn
}

func buildUpdateComandaStatus() {
	if UpdateComandaStatusClient != nil {
		return
	}
	repo := &repositories.UpdateComandaStatusRepository{Connection: resolveDB()}
	UpdateComandaStatusRepository = repo
	UpdateComandaStatusUseCase = &usecase.UpdateComandaStatus{Repository: repo}
	UpdateComandaStatusClient = &service.UpdateComandaStatusClient{UseCase: UpdateComandaStatusUseCase}
}

func ResolveUpdateComandaStatusContainer() iservice.IUpdateComandaStatus {
	buildUpdateComandaStatus()
	return UpdateComandaStatusClient
}
