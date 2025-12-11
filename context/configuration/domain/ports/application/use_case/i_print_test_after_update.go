package iusecase

import (
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IPrintTestAfterUpdate interface {
	Execute(resp *entities_main.Response[entities.PrinterIPUpdateResult]) (*entities_main.Response[entities.PrinterIPUpdateResult], error)
}
