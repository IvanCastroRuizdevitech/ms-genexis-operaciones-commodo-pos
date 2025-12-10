package irepositories

import (
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IUpdatePrinterIPRepository interface {
	Update(req *entities.PrinterIPUpdateRequest) (*entities_main.Response[entities.PrinterIPUpdateResult], error)
}
