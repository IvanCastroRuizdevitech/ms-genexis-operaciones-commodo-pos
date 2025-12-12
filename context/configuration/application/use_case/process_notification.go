package usecase

import (
	"log"
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	irepositories "ms-genexis-pos-operaciones/context/configuration/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type ProcessNotification struct {
	Repository irepositories.IProcessNotificationRepository
}

func (u *ProcessNotification) Execute(request *entities.ProcessNotificationRequest) (*entities_main.Response[entities.ProcessNotificationResult], error) {
	result, err := u.Repository.Process(request)
	if err != nil {
		log.Println("ProcessNotification usecase error:", err)
		return nil, err
	}

	return result, nil
}
