package domain_repositories

import "ms-genexis-pos-operaciones/domain/entities"

type IRecoverWacher interface {
	Consult(code string) (*entities.ParametersWatcher, error)
}
