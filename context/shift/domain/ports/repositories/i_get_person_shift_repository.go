package irepositories

import "ms-genexis-pos-operaciones/context/shift/domain/entities"

type IGetPersonShiftRepository interface {
	GetPersonShit(infoClient *entities.OpeningShiftRequest) (*entities.PersonShift, error)
}
