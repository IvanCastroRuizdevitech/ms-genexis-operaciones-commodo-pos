package irepositories

import (
	"ms-genexis-pos-operaciones/context/envelopes/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IGetTotalEnvelopesByJournalPromoterRepository interface {
	GetTotal(envelopes_request *entities.EnvelopesTotalRequest) (*entities_main.Response[entities.TotalEnvelopes], error)
}
