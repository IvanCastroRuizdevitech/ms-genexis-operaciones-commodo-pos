package irepositories

import (
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IValidateAdminPersonRepository interface {
	ValidateAdminPerson(request *entities.AdminValidationRequest) (*entities_main.Response[entities.AdminValidationResult], error)
}
