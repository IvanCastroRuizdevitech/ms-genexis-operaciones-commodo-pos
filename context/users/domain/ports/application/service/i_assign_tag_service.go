package iservice

import (
	"ms-genexis-pos-operaciones/context/users/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IAssignTag interface {
	Execute(request *entities.AssignTagRequest) (*entities_main.Response[entities.AssignTagResult], error)
}
