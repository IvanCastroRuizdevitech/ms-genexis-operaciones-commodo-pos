package service

import (
	"ms-genexis-pos-operaciones/context/envelopes/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/envelopes/domain/ports/application/use_case"
)

type TotalEnvelopesByJournalPromoterClient struct {
	GetTotalEnvelopes iusecase.IGetTotalEnvelopesByJournalPromoter
}

func (envelopes *TotalEnvelopesByJournalPromoterClient) Execute(envelopes_request *entities.EnvelopesTotalRequest) (*entities.ResponseEvelopesTotal, error) {
	db_response, err := envelopes.GetTotalEnvelopes.Execute(envelopes_request)
	if err != nil {
		return nil, err
	}

	return db_response, nil
}
