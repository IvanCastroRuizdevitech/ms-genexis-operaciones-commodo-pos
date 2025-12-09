package irepositories

import (
	entities_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IGetDispenserDetailsRepository interface {
	Get() (*entities_main.Response[entities_creditcustomers.DispenserDetailsFunctionResponse], error)
}
