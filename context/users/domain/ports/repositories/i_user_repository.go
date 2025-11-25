package irepositories

import (
	"ms-genexis-pos-operaciones/context/users/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IUserRepository interface {
	GetAll() (*entities_main.Response[[]entities.User], error)
	AssignTag(request *entities.AssignTagRequest) (*entities_main.Response[entities.AssignTagResult], error)
	GenerateAssignTagTransmissions(request *entities.AssignTagTransmissionsRequest) (*entities_main.Response[entities.AssignTagTransmissionsResult], error)
}
