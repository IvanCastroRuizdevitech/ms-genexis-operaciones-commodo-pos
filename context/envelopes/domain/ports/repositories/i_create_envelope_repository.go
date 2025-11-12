package irepositories

import (
	"ms-genexis-pos-operaciones/context/envelopes/domain/entities"
)

type ICreateEnvelopeRepository interface {
	Create(envelope_request *entities.EnvelopeRequest) (*entities.EnvelopeCreate, error)
}
