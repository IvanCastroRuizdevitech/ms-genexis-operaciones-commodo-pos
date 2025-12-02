package usecase

import (
	"log"

	entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
	irepositories "ms-genexis-pos-operaciones/context/reports/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetShiftSummary struct {
	Repository irepositories.IGetShiftSummaryRepository
}

func (u *GetShiftSummary) Execute(pos int, fechaInicio string, fechaFin string) (*entities_main.Response[entities_reports.ShiftSummary], error) {
	result, err := u.Repository.Get(pos, fechaInicio, fechaFin)
	if err != nil {
		log.Println("[GetShiftSummary usecase]", err)
		return nil, err
	}
	return result, nil
}
