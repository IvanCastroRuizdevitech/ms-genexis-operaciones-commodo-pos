package iservice

import (
	"ms-genexis-pos-operaciones/context/envelopes/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type ITotalEnvelopesByJournalPromoter interface {
	Execute(envelopes_request *entities.EnvelopesTotalRequest) (*entities_main.Response[entities.TotalEnvelopes], error)
}
