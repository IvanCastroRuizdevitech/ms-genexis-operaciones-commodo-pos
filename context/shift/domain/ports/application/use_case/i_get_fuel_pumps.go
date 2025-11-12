package iusecase

type IGetFuelPumps interface {
    Execute() ([]map[string]interface{}, error)
}

