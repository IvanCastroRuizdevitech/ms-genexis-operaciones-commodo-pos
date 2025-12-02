package service

import (
	iusecase "ms-genexis-pos-operaciones/context/sales/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type UpsertActiveDispenserFaceTransactionClient struct {
	UseCase iusecase.IUpsertActiveDispenserFaceTransaction
}

func (s *UpsertActiveDispenserFaceTransactionClient) Execute(surtidor int, cara int, codigo string, grado int, proveedorID int, montoMaximo float64, cantidadMaxima float64, trama map[string]any, promotorID int) (*entities_main.Response[map[string]any], error) {
	return s.UseCase.Execute(surtidor, cara, codigo, grado, proveedorID, montoMaximo, cantidadMaxima, trama, promotorID)
}
