package irepositories

import entities_main "ms-genexis-pos-operaciones/domain/entities"

type IUpsertActiveDispenserFaceTransactionRepository interface {
	Execute(surtidor int, cara int, codigo string, grado int, proveedorID int, montoMaximo float64, cantidadMaxima float64, trama map[string]any, promotorID int) (*entities_main.Response[map[string]any], error)
}
