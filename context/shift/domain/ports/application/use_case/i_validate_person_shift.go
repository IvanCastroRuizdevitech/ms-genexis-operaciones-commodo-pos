package iusecase

import "ms-genexis-pos-operaciones/context/shift/domain/entities"

type IvalidatePersonShift interface {
	Execute(shift_info *entities.OpeningShiftRequest) (*entities.PersonShift, error)
}
