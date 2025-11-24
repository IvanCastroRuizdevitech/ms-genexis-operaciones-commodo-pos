package iservice

import (
	"ms-genexis-pos-operaciones/context/users/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IUsersService interface {
	Execute() (*entities_main.Response[[]entities.User], error)
}

// IAssignTag is defined in i_assign_tag_service.go for clarity.
