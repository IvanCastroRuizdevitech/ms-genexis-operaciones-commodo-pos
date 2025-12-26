package usecase

import (
	"log"

	"ms-genexis-pos-operaciones/context/shift/domain/entities"
	irepositories "ms-genexis-pos-operaciones/context/shift/domain/ports/repositories"
)

type GetPersonByID struct {
	Repository irepositories.IGetPersonByIDRepository
}

func (u *GetPersonByID) Execute(personaID int64) (*entities.PersonByIDResponse, error) {
	result, err := u.Repository.GetPersonByID(personaID)
	if err != nil {
		log.Println("GetPersonByID: ", err)
		return nil, err
	}
	return result, nil
}
