package iusecase

import (
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IGetUnifiedConsecutives interface {
	Execute() (*entities_main.Response[entities.UnifiedConsecutivesResponse], error)
}
