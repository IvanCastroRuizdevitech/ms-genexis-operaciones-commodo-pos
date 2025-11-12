package irepositories

import "time"

type IGetDailyIncomeMeasurementsRepository interface {
    GetDailyIncomeMeasurements(start time.Time, end time.Time) ([]map[string]interface{}, error)
}

