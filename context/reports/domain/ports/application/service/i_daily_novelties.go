package iservice

import (
	entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IDailyNovelties interface {
	Execute(ano, mes, dia int) (*entities_main.Response[entities_reports.DailyNovelties], error)
}
