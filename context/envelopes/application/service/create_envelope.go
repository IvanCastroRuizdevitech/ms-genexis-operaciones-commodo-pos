package service

import (
	"ms-genexis-pos-operaciones/context/envelopes/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type CreateEnvelopeClient struct {
}

func (envelope *CreateEnvelopeClient) Execute(envelopes_request *entities.EnvelopeRequest) (*entities_main.Response[interface{}], error) {

	return nil, nil
}
