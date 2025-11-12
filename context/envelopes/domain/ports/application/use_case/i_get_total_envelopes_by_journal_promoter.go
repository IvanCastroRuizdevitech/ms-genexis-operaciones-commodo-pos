package iusecase

import "ms-genexis-pos-operaciones/context/envelopes/domain/entities"

type IGetTotalEnvelopesByJournalPromoter interface {
	Execute(envelopes_request *entities.EnvelopesTotalRequest) (*entities.ResponseEvelopesTotal, error)
}
