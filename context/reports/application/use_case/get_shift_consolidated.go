package usecase

import (
	"log"

	entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
	irepositories "ms-genexis-pos-operaciones/context/reports/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetShiftConsolidated struct {
	Repository irepositories.IGetShiftConsolidatedRepository
}

func (u *GetShiftConsolidated) Execute(fechaInicio string, fechaFin string) (*entities_main.Response[entities_reports.ShiftConsolidated], error) {
	result, err := u.Repository.Get(fechaInicio, fechaFin)
	if err != nil {
		log.Println("[GetShiftConsolidated usecase]", err)
		return nil, err
	}
	return result, nil
}
