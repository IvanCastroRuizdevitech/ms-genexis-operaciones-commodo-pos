package service

import (
	entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/reports/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type DailyNoveltiesClient struct {
	GetDailyNovelties iusecase.IGetDailyNovelties
}

func (s *DailyNoveltiesClient) Execute(ano, mes, dia int) (*entities_main.Response[entities_reports.DailyNovelties], error) {
	return s.GetDailyNovelties.Execute(ano, mes, dia)
}
