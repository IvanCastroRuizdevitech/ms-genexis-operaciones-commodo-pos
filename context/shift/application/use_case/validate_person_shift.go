package usecase

import (
	"log"
	"ms-genexis-pos-operaciones/context/shift/domain/entities"
	irepositories "ms-genexis-pos-operaciones/context/shift/domain/ports/repositories"
)

type ValidatePersonShift struct {
	ShiftRepository irepositories.IGetPersonShiftRepository
}

func (shift *ValidatePersonShift) Execute(shift_info *entities.OpeningShiftRequest) (*entities.PersonShift, error) {

	person, err := shift.ShiftRepository.GetPersonShit(shift_info)

	if err != nil {
		log.Println("ValidatePersonShift: ", err)
		return nil, err
	}

	return person, nil
}
