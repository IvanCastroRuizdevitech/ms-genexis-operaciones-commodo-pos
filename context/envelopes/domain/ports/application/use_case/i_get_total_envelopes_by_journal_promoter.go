package iusecase

import (
	"ms-genexis-pos-operaciones/context/envelopes/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IGetTotalEnvelopesByJournalPromoter interface {
	Execute(envelopes_request *entities.EnvelopesTotalRequest) (*entities_main.Response[entities.TotalEnvelopes], error)
}
