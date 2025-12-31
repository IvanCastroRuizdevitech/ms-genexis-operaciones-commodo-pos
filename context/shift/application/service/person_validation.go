package service

import (
	"net/http"

	"ms-genexis-pos-operaciones/context/shift/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/shift/domain/ports/application/use_case"
)

type PersonValidationClient struct {
	ValidatePerson iusecase.IPersonValidation
}

func (s *PersonValidationClient) ExecutePersonValidation(info *entities.PersonValidationRequest, requireAdmin bool) (*entities.PersonValidationResult, int, error) {
	result, err := s.ValidatePerson.Execute(info, requireAdmin)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	status := http.StatusOK
	if result.Message == "PARAMETROS_INVALIDOS" {
		status = http.StatusBadRequest
	} else if !result.Authenticated {
		status = http.StatusUnauthorized
	} else if !result.Success {
		status = http.StatusForbidden
	}

	return result, status, nil
}
