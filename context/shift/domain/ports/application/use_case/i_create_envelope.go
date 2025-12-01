package iusecase

import (
	"ms-genexis-pos-operaciones/context/shift/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type ICreateEnvelope interface {
	Execute(envelopeRequest *entities.EnvelopeRequest) (*entities_main.Response[entities.EnvelopeCreate], error)
}
