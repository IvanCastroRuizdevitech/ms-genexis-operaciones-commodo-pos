package irepositories

import (
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IGetUnifiedConsecutivesRepository interface {
	GetUnifiedConsecutives() (*entities_main.Response[entities.UnifiedConsecutivesResponse], error)
}
