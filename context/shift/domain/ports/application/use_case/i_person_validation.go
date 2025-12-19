package iusecase

import "ms-genexis-pos-operaciones/context/shift/domain/entities"

type IPersonValidation interface {
	Execute(info *entities.PersonValidationRequest, requireAdmin bool) (*entities.PersonValidationResult, error)
}
