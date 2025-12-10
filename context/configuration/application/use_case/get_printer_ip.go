package usecase

import (
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	irepositories "ms-genexis-pos-operaciones/context/configuration/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetPrinterIP struct {
	Repository irepositories.IGetPrinterIPRepository
}

func (u *GetPrinterIP) Execute() (*entities_main.Response[entities.PrinterIPResult], error) {
	return u.Repository.Get()
}
