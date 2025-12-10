package usecase

import (
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	irepositories "ms-genexis-pos-operaciones/context/configuration/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type UpdatePrinterIP struct {
	Repository irepositories.IUpdatePrinterIPRepository
}

func (u *UpdatePrinterIP) Execute(req *entities.PrinterIPUpdateRequest) (*entities_main.Response[entities.PrinterIPUpdateResult], error) {
	return u.Repository.Update(req)
}
