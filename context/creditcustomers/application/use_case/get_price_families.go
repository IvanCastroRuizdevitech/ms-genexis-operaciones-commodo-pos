package use_case

import (
	entities_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	irepositories "ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetPriceFamilies struct {
	Repository irepositories.IGetPriceFamiliesRepository
}

func (u *GetPriceFamilies) Execute() (*entities_main.Response[[]entities_creditcustomers.PriceFamily], error) {
	return u.Repository.Get()
}
