package usecase

import (
	"log"
	entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/reports/domain/ports/application/use_case"
	irepositories "ms-genexis-pos-operaciones/context/reports/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

var _ iusecase.IGetDailyNovelties = (*GetDailyNovelties)(nil)

type GetDailyNovelties struct {
	Repository irepositories.IGetDailyNoveltiesRepository
}

func (u *GetDailyNovelties) Execute(ano, mes, dia int) (*entities_main.Response[entities_reports.DailyNovelties], error) {
	result, err := u.Repository.GetNovelties(ano, mes, dia)
	if err != nil {
		log.Println("GetDailyNovelties error:", err)
		return nil, err
	}

	return result, nil
}
