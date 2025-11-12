package iservice

import "ms-genexis-pos-operaciones/context/shift/domain/entities"

type IOpeningShift interface {
	ExecuteOpeningShift(shift_info *entities.OpeningShiftRequest) (*entities.ResponseShift, error)
}
