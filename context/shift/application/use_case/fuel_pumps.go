package usecase

import (
    "ms-genexis-pos-operaciones/context/shift/domain/entities"
    irepositories "ms-genexis-pos-operaciones/context/shift/domain/ports/repositories"
)

type GetFuelPumps struct {
    Repository irepositories.IGetFuelPumpsRepository
}

func (u *GetFuelPumps) Execute(q *entities.FuelPumpsRequest) ([]map[string]interface{}, error) {
    return u.Repository.GetFuelPumps(q.TurnoId, q.EquiposId)
}

