package irepositories

import (
    "ms-genexis-pos-operaciones/context/home/domain/entities"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type ILoadErrorNotificationRepository interface {
    Load() (*entities_main.Response[[]entities.ErrorNotification], error)
}
