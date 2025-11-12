package usecase

import (
    "ms-genexis-pos-operaciones/context/shift/domain/entities"
    irepositories "ms-genexis-pos-operaciones/context/shift/domain/ports/repositories"
)

type GetDailyIncomeMeasurements struct {
    Repository irepositories.IGetDailyIncomeMeasurementsRepository
}

func (u *GetDailyIncomeMeasurements) Execute(q *entities.DailyIncomeMeasurementsQuery) ([]map[string]interface{}, error) {
    return u.Repository.GetDailyIncomeMeasurements(q.FechaInicio, q.FechaFin)
}

