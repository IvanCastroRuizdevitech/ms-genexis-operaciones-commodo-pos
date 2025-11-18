package irepositories

type IGetFuelPumpsRepository interface {
    GetFuelPumps() ([]map[string]interface{}, error)
}
