package irepositories

type IGetFuelPumpsRepository interface {
    GetFuelPumps(turnoId int64, equiposId int64) ([]map[string]interface{}, error)
}

