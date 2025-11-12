package iusecase

import "ms-genexis-pos-operaciones/context/shift/domain/entities"

type IOpeningShift interface {
	Execute(shift_info *entities.OpeningShiftRequest, person_id int) (*entities.OpeningShiftHttpResponse, error)
}
