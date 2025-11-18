package service

import (
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    iusecase "ms-genexis-pos-operaciones/context/sales/domain/ports/application/use_case"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetPendingSaleDatafonoClient struct {
    UseCase iusecase.IGetPendingSaleDatafono
}

func (s *GetPendingSaleDatafonoClient) Execute(transactionId int) (*entities_main.Response[entities_sales.PendingSaleDatafono], error) {
    return s.UseCase.Execute(transactionId)
}

