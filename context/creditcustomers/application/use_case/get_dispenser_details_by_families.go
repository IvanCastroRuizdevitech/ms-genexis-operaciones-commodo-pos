package use_case

import (
	entities_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	irepositories "ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetDispenserDetailsByFamilies struct {
	Repository irepositories.IGetDispenserDetailsByFamiliesRepository
}

func (u *GetDispenserDetailsByFamilies) Execute(request *entities_creditcustomers.DispenserDetailsByFamiliesRequest) (*entities_main.Response[entities_creditcustomers.DispenserDetailsFunctionResponse], error) {
	return u.Repository.Get(request)
}
