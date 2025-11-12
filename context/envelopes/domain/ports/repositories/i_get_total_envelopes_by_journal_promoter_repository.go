package irepositories

import "ms-genexis-pos-operaciones/context/envelopes/domain/entities"

type IGetTotalEnvelopesByJournalPromoterRepository interface {
	GetTotal(envelopes_request *entities.EnvelopesTotalRequest) (*entities.ResponseEvelopesTotal, error)
}
