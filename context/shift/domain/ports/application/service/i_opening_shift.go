package iservice

import "ms-genexis-pos-operaciones/context/shift/domain/entities"

type IOpeningShift interface {
	ExecuteOpeningShift(shiftInfo *entities.OpeningShiftRequest) (*entities.ResponseShift, error)
}
