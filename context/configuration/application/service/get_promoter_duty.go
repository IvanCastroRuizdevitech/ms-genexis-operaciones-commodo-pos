package service

import (
    "ms-genexis-pos-operaciones/context/configuration/domain/entities"
    iusecase "ms-genexis-pos-operaciones/context/configuration/domain/ports/application/use_case"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetPromoterDutyClient struct {
    GetPromoterDuty iusecase.IGetPromoterDuty
}

func (c *GetPromoterDutyClient) Execute() (*entities_main.Response[[]entities.PromoterDuty], error) {
    dbResponse, err := c.GetPromoterDuty.Execute()
    if err != nil {
        return nil, err
    }
    return dbResponse, nil
}

