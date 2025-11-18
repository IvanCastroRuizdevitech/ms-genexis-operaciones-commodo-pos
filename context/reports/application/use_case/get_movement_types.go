package usecase

import (
	"log"
	entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/reports/domain/ports/application/use_case"
	irepositories "ms-genexis-pos-operaciones/context/reports/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

var _ iusecase.IGetMovementTypes = (*GetMovementTypes)(nil)

type GetMovementTypes struct {
	Repository irepositories.IGetMovementTypesRepository
}

func (u *GetMovementTypes) Execute() (*entities_main.Response[[]entities_reports.MovementType], error) {
	result, err := u.Repository.GetAll()
	if err != nil {
		log.Println("GetMovementTypes error:", err)
		return nil, err
	}

	return result, nil
}
