package container_creditcustomers

import (
	"ms-genexis-pos-operaciones/context/creditcustomers/application/service"
	usecase "ms-genexis-pos-operaciones/context/creditcustomers/application/use_case"
	iservice "ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/application/service"
	iusecase "ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/application/use_case"
	irepositories "ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/repositories"
	repositories "ms-genexis-pos-operaciones/context/creditcustomers/infrastructure"
	presentation_container "ms-genexis-pos-operaciones/presentation/container"
)

var GetDispenserDetailsRepository irepositories.IGetDispenserDetailsRepository
var GetDispenserDetailsUseCase iusecase.IGetDispenserDetails
var GetDispenserDetailsClient iservice.IGetDispenserDetailsService

func buildGetDispenserDetails() {
	if GetDispenserDetailsClient != nil {
		return
	}
	GetDispenserDetailsRepository = &repositories.GetDispenserDetailsRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}
	GetDispenserDetailsUseCase = &usecase.GetDispenserDetails{Repository: GetDispenserDetailsRepository}
	GetDispenserDetailsClient = &service.GetDispenserDetailsService{UseCase: GetDispenserDetailsUseCase}
}

func ResolveGetDispenserDetailsContainer() iservice.IGetDispenserDetailsService {
	buildGetDispenserDetails()
	return GetDispenserDetailsClient
}
