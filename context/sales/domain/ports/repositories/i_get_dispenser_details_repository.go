package irepositories

import (
	entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IGetDispenserDetailsRepository interface {
	Get() (*entities_main.Response[[]entities_sales.DispenserDetail], error)
}
