package service

import (
	"log"
	"ms-genexis-pos-operaciones/context/shift/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/shift/domain/ports/application/use_case"
	"time"
)

type OpeningShiftClient struct {
	ValidatePerson iusecase.IvalidatePersonShift
	OpeningShift   iusecase.IOpeningShift
}

func (shift *OpeningShiftClient) ExecuteOpeningShift(shift_info *entities.OpeningShiftRequest) (*entities.ResponseShift, error) {
	person_identified, err := shift.ValidatePerson.Execute(shift_info)
	if err != nil {
		return nil, err
	}

	log.Println("Persona identificada => ", person_identified)
	response_shift, err := shift.OpeningShift.Execute(shift_info, person_identified.Id)
	if err != nil {
		return nil, err
	}

	response := &entities.ResponseShift{
		Status:      response_shift.Estado,
		Message:     response_shift.Mensaje,
		ProcessDate: time.Now().Format("2006-01-02 15:04:05"),
		Data:        response_shift.Data,
	}
	return response, nil
}
