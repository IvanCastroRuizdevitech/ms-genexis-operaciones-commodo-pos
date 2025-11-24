package service

import (
	"ms-genexis-pos-operaciones/context/home/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/home/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type PendingTransmissionsProcessor struct {
	ProcessPendingTransmissions iusecase.IProcessPendingTransmissions
}

func (s *PendingTransmissionsProcessor) Execute() (*entities_main.Response[entities.TransmissionProcessSummary], error) {
	return s.ProcessPendingTransmissions.Execute()
}
