package usecase

import (
    "log"
    "ms-genexis-pos-operaciones/context/home/domain/entities"
    irepositories "ms-genexis-pos-operaciones/context/home/domain/ports/repositories"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type LoadErrorNotification struct {
    Repository irepositories.ILoadErrorNotificationRepository
}

func (u *LoadErrorNotification) Execute() (*entities_main.Response[[]entities.ErrorNotification], error) {
    result, err := u.Repository.Load()
    if err != nil {
        log.Println("[LoadErrorNotification][Execute]", err)
        return nil, err
    }
    return result, nil
}
