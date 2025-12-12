package iservice

import (
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IValidateAdminPerson interface {
	Execute(request *entities.AdminValidationRequest) (*entities_main.Response[entities.AdminValidationResult], error)
}
