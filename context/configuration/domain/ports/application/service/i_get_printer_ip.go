package iservice

import (
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IGetPrinterIP interface {
	Execute() (*entities_main.Response[entities.PrinterIPResult], error)
}
