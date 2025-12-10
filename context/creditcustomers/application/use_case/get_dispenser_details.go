package use_case

import (
	entities_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	irepositories "ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetDispenserDetails struct {
	Repository irepositories.IGetDispenserDetailsRepository
}

func (u *GetDispenserDetails) Execute() (*entities_main.Response[entities_creditcustomers.DispenserDetailsFunctionResponse], error) {
	return u.Repository.Get()
}
