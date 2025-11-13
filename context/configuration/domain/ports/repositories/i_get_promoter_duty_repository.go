package irepositories

import (
    "ms-genexis-pos-operaciones/context/configuration/domain/entities"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IGetPromoterDutyRepository interface {
    GetPromoterDuty() (*entities_main.Response[[]entities.PromoterDuty], error)
}

