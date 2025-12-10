package irepositories

import (
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IGetPrinterIPRepository interface {
	Get() (*entities_main.Response[entities.PrinterIPResult], error)
}
