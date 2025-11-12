package irepositories

import "ms-genexis-pos-operaciones/context/shift/domain/entities"

type ISendOpeningShiftRepositoryHttp interface {
	Send(url string, body []byte, header *map[string]string) (*entities.OpeningShiftHttpResponse, error)
}
