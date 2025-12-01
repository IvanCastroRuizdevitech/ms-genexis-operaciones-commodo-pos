package usecase

import (
	"log"
	"ms-genexis-pos-operaciones/context/shift/domain/entities"
	irepositories "ms-genexis-pos-operaciones/context/shift/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	"time"
)

type CreateEnvelope struct {
	Repository irepositories.ICreateEnvelopeRepository
}

func (uc *CreateEnvelope) Execute(envelopeRequest *entities.EnvelopeRequest) (*entities_main.Response[entities.EnvelopeCreate], error) {
	result, err := uc.Repository.Create(envelopeRequest)
	if err != nil {
		log.Println("CreateEnvelope: ", err)
		return nil, err
	}

	success := entities_main.NewSuccessResponse(
		200,
		"Sobre creado",
		time.Now().Format("2006-01-02 15:04:05"),
		result,
	)

	return &success, nil
}
