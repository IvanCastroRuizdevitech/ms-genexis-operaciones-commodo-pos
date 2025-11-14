package service

import (
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    iusecase "ms-genexis-pos-operaciones/context/sales/domain/ports/application/use_case"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetUnresolvedSaleAttributesClient struct {
    UseCase iusecase.IGetUnresolvedSaleAttributes
}

func (s *GetUnresolvedSaleAttributesClient) Execute(movementId int) (*entities_main.Response[entities_sales.UnresolvedSaleAttributes], error) {
    return s.UseCase.Execute(movementId)
}

