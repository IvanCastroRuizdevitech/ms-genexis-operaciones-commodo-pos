package service

import (
	"net/http"
	"time"

	"ms-genexis-pos-operaciones/context/shift/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/shift/domain/ports/application/use_case"
)

type PersonValidationClient struct {
	ValidatePerson iusecase.IPersonValidation
}

func (s *PersonValidationClient) ExecutePersonValidation(info *entities.PersonValidationRequest, requireAdmin bool) (*entities.ResponseShift, error) {
	result, err := s.ValidatePerson.Execute(info, requireAdmin)
	if err != nil {
		return nil, err
	}

	status := http.StatusOK
	if result.Message == "PARAMETROS_INVALIDOS" {
		status = http.StatusBadRequest
	} else if !result.Authenticated {
		status = http.StatusUnauthorized
	} else if !result.Success {
		status = http.StatusForbidden
	}

	return &entities.ResponseShift{
		Status:      status,
		Message:     result.Message,
		ProcessDate: time.Now().Format("2006-01-02 15:04:05"),
		Data:        result,
	}, nil
}
