package usecase

import (
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	irepositories "ms-genexis-pos-operaciones/context/configuration/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetUnifiedConsecutives struct {
	Repository irepositories.IGetUnifiedConsecutivesRepository
}

func (u *GetUnifiedConsecutives) Execute() (*entities_main.Response[entities.UnifiedConsecutivesResponse], error) {
	return u.Repository.GetUnifiedConsecutives()
}
