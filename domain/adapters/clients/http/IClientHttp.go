package domain_adapters_clients_http

import "ms-genexis-pos-operaciones/domain/entities"

type IClientHttp interface {
	Send(method string, url string, message *entities.HttpRequest) (*entities.HttpResponse, error)
}
