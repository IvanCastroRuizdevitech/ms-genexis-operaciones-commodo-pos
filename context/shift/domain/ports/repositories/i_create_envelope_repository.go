package irepositories

import "ms-genexis-pos-operaciones/context/shift/domain/entities"

type ICreateEnvelopeRepository interface {
	Create(envelopeRequest *entities.EnvelopeRequest) (*entities.EnvelopeCreate, error)
}
