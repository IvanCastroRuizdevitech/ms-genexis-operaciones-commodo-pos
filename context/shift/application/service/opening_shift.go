package service

import "ms-genexis-pos-operaciones/context/shift/domain/entities"

type OpeningShiftClient struct {
}

func (shift *OpeningShiftClient) ExecuteOpeningShift(shiftInfo *entities.OpeningShiftRequest) (*entities.ResponseShift, error) {
	return nil, nil
}
