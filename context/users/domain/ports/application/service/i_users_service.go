package iservice

import (
	"ms-genexis-pos-operaciones/context/users/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IUsersService interface {
	Execute() (*entities_main.Response[[]entities.User], error)
}
