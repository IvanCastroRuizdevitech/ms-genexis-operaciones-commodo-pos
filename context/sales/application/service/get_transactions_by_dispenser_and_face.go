package service

import (
	iusecase "ms-genexis-pos-operaciones/context/sales/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetTransactionsByDispenserAndFaceClient struct {
	UseCase iusecase.IGetTransactionsByDispenserAndFace
}

func (s *GetTransactionsByDispenserAndFaceClient) Execute(dispenserID int, face int) (*entities_main.Response[[]map[string]interface{}], error) {
	return s.UseCase.Execute(dispenserID, face)
}
