package iservice

import (
	entities_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IGetDispenserDetailsService interface {
	Execute() (*entities_main.Response[entities_creditcustomers.DispenserDetailsFunctionResponse], error)
}
