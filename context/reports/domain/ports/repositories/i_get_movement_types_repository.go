package irepositories

import (
	entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IGetMovementTypesRepository interface {
	GetAll() (*entities_main.Response[[]entities_reports.MovementType], error)
}
