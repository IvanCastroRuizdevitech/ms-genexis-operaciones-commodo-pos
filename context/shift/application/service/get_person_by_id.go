package service

import (
	"ms-genexis-pos-operaciones/context/shift/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/shift/domain/ports/application/use_case"
)

type GetPersonByIDClient struct {
	GetPersonByID iusecase.IGetPersonByID
}

func (s *GetPersonByIDClient) ExecutePersonByID(personaID int64) (*entities.PersonByIDResponse, error) {
	return s.GetPersonByID.Execute(personaID)
}
