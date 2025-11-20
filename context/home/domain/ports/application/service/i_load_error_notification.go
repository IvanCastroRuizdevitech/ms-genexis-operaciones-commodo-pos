package iservice

import (
    "ms-genexis-pos-operaciones/context/home/domain/entities"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type ILoadErrorNotification interface {
    Execute() (*entities_main.Response[[]entities.ErrorNotification], error)
}
