package iusecase

import "ms-genexis-pos-operaciones/context/shift/domain/entities"

type IGetFuelPumps interface {
    Execute(q *entities.FuelPumpsRequest) ([]map[string]interface{}, error)
}

