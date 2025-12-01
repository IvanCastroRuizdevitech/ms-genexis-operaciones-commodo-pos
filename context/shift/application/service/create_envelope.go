package service

import (
	"log"
	"ms-genexis-pos-operaciones/context/shift/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/shift/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type CreateEnvelopeClient struct {
	CreateEnvelope iusecase.ICreateEnvelope
}

func (s *CreateEnvelopeClient) Execute(envelopeRequest *entities.EnvelopeRequest) (*entities_main.Response[entities.EnvelopeCreate], error) {
	result, err := s.CreateEnvelope.Execute(envelopeRequest)
	if err != nil {
		log.Println("CreateEnvelopeClient: ", err)
		return nil, err
	}
	return result, nil
}
