package usecase

import (
	"log"
	"ms-genexis-pos-operaciones/context/envelopes/domain/entities"
	irepositories "ms-genexis-pos-operaciones/context/envelopes/domain/ports/repositories"
)

type GetTotalEnvelopesByJournalPromoter struct {
	GetTotalEnvelopes irepositories.IGetTotalEnvelopesByJournalPromoterRepository
}

func (envelopes *GetTotalEnvelopesByJournalPromoter) Execute(envelopes_request *entities.EnvelopesTotalRequest) (*entities.ResponseEvelopesTotal, error) {

	result, err := envelopes.GetTotalEnvelopes.GetTotal(envelopes_request)

	if err != nil {
		log.Println("GetTotalEnvelopesByJournalPromoter: ", err)
		return nil, err
	}

	return result, nil
}
