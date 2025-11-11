package presentation_container

import (
	"log"
	domain_adapters_clients_db "ms-genexis-pos-operaciones/domain/adapters/clients/db"
	domain_adapters_clients_http "ms-genexis-pos-operaciones/domain/adapters/clients/http"
	domain_repositories "ms-genexis-pos-operaciones/domain/repositories/db"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
	infrastructure_http_client "ms-genexis-pos-operaciones/infrastructure/http/client"
)

// SERVICES

// USES CASES

// REPOSITORIES

// GENERALS
var _ domain_repositories.IRecoverWacher

var client_http domain_adapters_clients_http.IClientHttp
var client_db domain_adapters_clients_db.IClientDB

var err error

func InitContainer() error {

	//CLIENTS
	client_db, err = infrastructure_db_client.InitClient("")
	if err != nil {
		log.Fatal("[InitContainer] - Error init client_db", err)
		return err
	}

	client_http, err = infrastructure_http_client.InitClient()
	if err != nil {
		log.Fatal("[InitContainer] - Error init client_http", err)
		return err
	}

	// REPOSITORIES

	// USES CASES

	// SERVICE

	return nil
}
