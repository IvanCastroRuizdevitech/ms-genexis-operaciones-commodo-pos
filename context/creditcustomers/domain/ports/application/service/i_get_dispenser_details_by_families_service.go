package iservice

import (
	entities_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IGetDispenserDetailsByFamiliesService interface {
	Execute(request *entities_creditcustomers.DispenserDetailsByFamiliesRequest) (*entities_main.Response[entities_creditcustomers.DispenserDetailsFunctionResponse], error)
}
