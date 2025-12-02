package iusecase

import (
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IGetTransactionsByDispenserAndFace interface {
	Execute(dispenserID int, face int) (*entities_main.Response[[]map[string]interface{}], error)
}
