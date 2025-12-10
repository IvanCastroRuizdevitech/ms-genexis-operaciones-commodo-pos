package iusecase

import (
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IUpdatePrinterIP interface {
	Execute(req *entities.PrinterIPUpdateRequest) (*entities_main.Response[entities.PrinterIPUpdateResult], error)
}
