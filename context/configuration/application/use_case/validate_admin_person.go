package usecase

import (
	"log"
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/configuration/domain/ports/application/use_case"
	irepositories "ms-genexis-pos-operaciones/context/configuration/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type ValidateAdminPerson struct {
	Repository irepositories.IValidateAdminPersonRepository
}

func (u *ValidateAdminPerson) Execute(request *entities.AdminValidationRequest) (*entities_main.Response[entities.AdminValidationResult], error) {
	result, err := u.Repository.ValidateAdminPerson(request)
	if err != nil {
		log.Println("ValidateAdminPerson: ", err)
		return nil, err
	}
	return result, nil
}

var _ iusecase.IValidateAdminPerson = (*ValidateAdminPerson)(nil)
