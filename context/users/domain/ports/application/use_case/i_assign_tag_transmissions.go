package iusecase

import (
	"ms-genexis-pos-operaciones/context/users/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IAssignTagTransmissions interface {
	Execute(request *entities.AssignTagTransmissionsRequest) (*entities_main.Response[entities.AssignTagTransmissionsResult], error)
}
