package usecase

import (
	"log"
	"ms-genexis-pos-operaciones/context/shift/domain/entities"
	irepositories "ms-genexis-pos-operaciones/context/shift/domain/ports/repositories"
)

type PersonValidation struct {
	Repository irepositories.IValidatePersonRepository
}

func (p *PersonValidation) Execute(info *entities.PersonValidationRequest, requireAdmin bool) (*entities.PersonValidationResult, error) {
	result, err := p.Repository.ValidatePerson(info, requireAdmin)
	if err != nil {
		log.Println("PersonValidation: ", err)
		return nil, err
	}

	return result, nil
}
