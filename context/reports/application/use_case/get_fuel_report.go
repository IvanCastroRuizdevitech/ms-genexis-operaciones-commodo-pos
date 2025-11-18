package usecase

import (
	"log"
	entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/reports/domain/ports/application/use_case"
	irepositories "ms-genexis-pos-operaciones/context/reports/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

var _ iusecase.IGetFuelReport = (*GetFuelReport)(nil)

type GetFuelReport struct {
	Repository irepositories.IGetFuelReportRepository
}

func (u *GetFuelReport) Execute(fecha string) (*entities_main.Response[entities_reports.FuelReport], error) {
	result, err := u.Repository.GetReport(fecha)
	if err != nil {
		log.Println("GetFuelReport error:", err)
		return nil, err
	}
	return result, nil
}
