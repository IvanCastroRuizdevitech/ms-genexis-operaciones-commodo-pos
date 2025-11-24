package service

import (
	"ms-genexis-pos-operaciones/context/shift/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/shift/domain/ports/application/use_case"
	"time"
)

type PersonValidationClient struct {
	ValidatePerson iusecase.IPersonValidation
}

func (s *PersonValidationClient) ExecutePersonValidation(info *entities.PersonValidationRequest) (*entities.ResponseShift, error) {
	person, err := s.ValidatePerson.Execute(info)
	if err != nil {
		return nil, err
	}

	return &entities.ResponseShift{
		Status:      200,
		Message:     "ok",
		ProcessDate: time.Now().Format("2006-01-02 15:04:05"),
		Data:        person,
	}, nil
}
