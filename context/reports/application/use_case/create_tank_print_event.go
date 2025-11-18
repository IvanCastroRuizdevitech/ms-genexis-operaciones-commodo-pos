package usecase

import (
	"log"
	entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/reports/domain/ports/application/use_case"
	irepositories "ms-genexis-pos-operaciones/context/reports/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

var _ iusecase.ICreateTankPrintEvent = (*CreateTankPrintEvent)(nil)

type CreateTankPrintEvent struct {
	Repository irepositories.ICreateTankPrintEventRepository
}

func (u *CreateTankPrintEvent) Execute(tankIDs []int) (*entities_main.Response[entities_reports.TankPrintEventResult], error) {
	result, err := u.Repository.CreateEvent(tankIDs)
	if err != nil {
		log.Println("CreateTankPrintEvent error:", err)
		return nil, err
	}

	return result, nil
}
