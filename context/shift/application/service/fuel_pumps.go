package service

import (
    "time"
    "ms-genexis-pos-operaciones/context/shift/domain/entities"
    iusecase "ms-genexis-pos-operaciones/context/shift/domain/ports/application/use_case"
)

type FuelPumpsClient struct {
    GetFuelPumps iusecase.IGetFuelPumps
}

func (s *FuelPumpsClient) ExecuteFuelPumps(q *entities.FuelPumpsRequest) (*entities.ResponseShift, error) {
    data, err := s.GetFuelPumps.Execute(q)
    if err != nil {
        return nil, err
    }
    return &entities.ResponseShift{
        Status:      200,
        Message:     "ok",
        ProcessDate: time.Now().Format("2006-01-02 15:04:05"),
        Data:        data,
    }, nil
}

