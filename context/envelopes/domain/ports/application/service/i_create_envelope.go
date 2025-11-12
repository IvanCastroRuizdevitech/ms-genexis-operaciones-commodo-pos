package iservice

import (
	"ms-genexis-pos-operaciones/context/envelopes/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type ICreateEnvelope interface {
	Execute(envelopes_request *entities.EnvelopeRequest) (*entities_main.Response[interface{}], error)
}
