package iusecase

import "ms-genexis-pos-operaciones/context/shift/domain/entities"

type IGetPersonByID interface {
	Execute(personaID int64) (*entities.PersonByIDResponse, error)
}
