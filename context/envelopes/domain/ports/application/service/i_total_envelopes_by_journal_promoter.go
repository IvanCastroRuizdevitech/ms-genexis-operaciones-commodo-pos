package iservice

import "ms-genexis-pos-operaciones/context/envelopes/domain/entities"

type ITotalEnvelopesByJournalPromoter interface {
	Execute(envelopes_request *entities.EnvelopesTotalRequest) (*entities.ResponseEvelopesTotal, error)
}
