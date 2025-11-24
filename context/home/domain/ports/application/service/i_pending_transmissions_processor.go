package iservice

import (
	"ms-genexis-pos-operaciones/context/home/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IPendingTransmissionsProcessor interface {
	Execute() (*entities_main.Response[entities.TransmissionProcessSummary], error)
}
