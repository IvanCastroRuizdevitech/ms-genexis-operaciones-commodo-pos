package container_envelopes

import (
	"ms-genexis-pos-operaciones/context/envelopes/application/service"
	usecase "ms-genexis-pos-operaciones/context/envelopes/application/use_case"
	iservice "ms-genexis-pos-operaciones/context/envelopes/domain/ports/application/service"
	iusecase "ms-genexis-pos-operaciones/context/envelopes/domain/ports/application/use_case"
	irepositories "ms-genexis-pos-operaciones/context/envelopes/domain/ports/repositories"
	repositories "ms-genexis-pos-operaciones/context/envelopes/infrastructure"
	presentation_container "ms-genexis-pos-operaciones/presentation/container"
)

// REPOSITORIES DB
var GetTotalEnvelopesByJournalPromoterRepository irepositories.IGetTotalEnvelopesByJournalPromoterRepository

// REPOSITORIES HTTPP

// USECASE
var GetTotalEnvelopesByJournalPromoterUseCase iusecase.IGetTotalEnvelopesByJournalPromoter

// SERVICE
var TotalEnvelopesByJournalPromoterClient iservice.ITotalEnvelopesByJournalPromoter

func initializes() {
	GetTotalEnvelopesByJournalPromoterRepository = &repositories.GetTotalEnvelopesByJournalPromoterRepository{Connection: presentation_container.ResolveDatabaseConnectionToLecWithPgx()}
	GetTotalEnvelopesByJournalPromoterUseCase = &usecase.GetTotalEnvelopesByJournalPromoter{GetTotalEnvelopes: GetTotalEnvelopesByJournalPromoterRepository}
	TotalEnvelopesByJournalPromoterClient = &service.TotalEnvelopesByJournalPromoterClient{GetTotalEnvelopes: GetTotalEnvelopesByJournalPromoterUseCase}
}

func ResolveEnvelopesContainer() iservice.ITotalEnvelopesByJournalPromoter {

	if TotalEnvelopesByJournalPromoterClient == nil {
		initializes()
	}

	return TotalEnvelopesByJournalPromoterClient
}
