package service

import (
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    iusecase "ms-genexis-pos-operaciones/context/sales/domain/ports/application/use_case"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type ReprintSaleClient struct {
    UseCase iusecase.IReprintSale
}

func (s *ReprintSaleClient) Execute(movementId int) (*entities_main.Response[entities_sales.ReprintSaleResult], error) {
    return s.UseCase.Execute(movementId)
}

