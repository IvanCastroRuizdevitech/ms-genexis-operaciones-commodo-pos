package service

import (
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	iservice "ms-genexis-pos-operaciones/context/configuration/domain/ports/application/service"
	iusecase "ms-genexis-pos-operaciones/context/configuration/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type ValidateAdminPersonClient struct {
	UseCase iusecase.IValidateAdminPerson
}

func (s *ValidateAdminPersonClient) Execute(request *entities.AdminValidationRequest) (*entities_main.Response[entities.AdminValidationResult], error) {
	response, err := s.UseCase.Execute(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

var _ iservice.IValidateAdminPerson = (*ValidateAdminPersonClient)(nil)
