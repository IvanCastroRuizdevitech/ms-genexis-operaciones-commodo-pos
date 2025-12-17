package use_case

import (
	entities_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	irepositories "ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetIdentifierTypes struct {
	Repository irepositories.IGetIdentifierTypesRepository
}

func (u *GetIdentifierTypes) Execute() (*entities_main.Response[[]entities_creditcustomers.IdentifierType], error) {
	return u.Repository.Get()
}
