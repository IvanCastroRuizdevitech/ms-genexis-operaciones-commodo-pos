package service

import (
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/configuration/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetUnifiedConsecutivesClient struct {
	UseCase iusecase.IGetUnifiedConsecutives
}

func (s *GetUnifiedConsecutivesClient) Execute() (*entities_main.Response[entities.UnifiedConsecutivesResponse], error) {
	return s.UseCase.Execute()
}
