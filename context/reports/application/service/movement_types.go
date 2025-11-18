package service

import (
	entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/reports/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type MovementTypesClient struct {
	GetMovementTypes iusecase.IGetMovementTypes
}

func (s *MovementTypesClient) Execute() (*entities_main.Response[[]entities_reports.MovementType], error) {
	return s.GetMovementTypes.Execute()
}
