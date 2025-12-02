package irepositories

import (
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IGetTransactionsByDispenserAndFaceRepository interface {
	Get(dispenserID int, face int) (*entities_main.Response[[]map[string]interface{}], error)
}
