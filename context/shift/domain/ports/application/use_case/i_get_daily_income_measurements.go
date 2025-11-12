package iusecase

import (
    "ms-genexis-pos-operaciones/context/shift/domain/entities"
)

type IGetDailyIncomeMeasurements interface {
    Execute(q *entities.DailyIncomeMeasurementsQuery) ([]map[string]interface{}, error)
}

