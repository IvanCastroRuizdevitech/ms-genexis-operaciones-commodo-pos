package usecase

import (
	"log"
	"ms-genexis-pos-operaciones/context/shift/domain/entities"
	irepositories "ms-genexis-pos-operaciones/context/shift/domain/ports/repositories"
)

type PersonValidation struct {
	Repository irepositories.IValidatePersonRepository
}

func (p *PersonValidation) Execute(info *entities.PersonValidationRequest) (*entities.PersonShift, error) {
	person, err := p.Repository.ValidatePerson(info)
	if err != nil {
		log.Println("PersonValidation: ", err)
		return nil, err
	}

	return person, nil
}
