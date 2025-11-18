package usecase

import (
    irepositories "ms-genexis-pos-operaciones/context/shift/domain/ports/repositories"
)

type GetFuelPumps struct {
    Repository irepositories.IGetFuelPumpsRepository
}

func (u *GetFuelPumps) Execute() ([]map[string]interface{}, error) {
    return u.Repository.GetFuelPumps()
}
