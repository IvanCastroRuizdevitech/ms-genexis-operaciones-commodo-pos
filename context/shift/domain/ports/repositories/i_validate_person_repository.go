package irepositories

import "ms-genexis-pos-operaciones/context/shift/domain/entities"

type IValidatePersonRepository interface {
	ValidatePerson(info *entities.PersonValidationRequest) (*entities.PersonShift, error)
}
