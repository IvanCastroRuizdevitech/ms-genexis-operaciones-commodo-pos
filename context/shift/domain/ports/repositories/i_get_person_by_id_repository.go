package irepositories

import "ms-genexis-pos-operaciones/context/shift/domain/entities"

type IGetPersonByIDRepository interface {
	GetPersonByID(personaID int64) (*entities.PersonByIDResponse, error)
}
