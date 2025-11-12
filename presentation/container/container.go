package presentation_container

import (
	"ms-genexis-pos-operaciones/domain/constants"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
	infrastructure_db_client_postgres_drivers "ms-genexis-pos-operaciones/infrastructure/db/client/postgres/drivers"
	infrastructure_external_nethttp "ms-genexis-pos-operaciones/infrastructure/externals/externalhttp"
	infrastructure_external_nethttp_nethttp "ms-genexis-pos-operaciones/infrastructure/externals/externalhttp/nethttp"
)

var DatabaseConnectionToLecWithPgx infrastructure_db_client.DatabaseConnectionInterface

func ResolveDatabaseConnectionToLecWithPgx() infrastructure_db_client.DatabaseConnectionInterface {
	if DatabaseConnectionToLecWithPgx != nil {
		return DatabaseConnectionToLecWithPgx
	}
	return &infrastructure_db_client_postgres_drivers.ConfigConnectionPgx{
		UrlToConnect: constants.DB_CON,
	}
}

var ClientHttpWithNet infrastructure_external_nethttp.ClientHTTPInterface

func ResolveClientHttpWithNet() infrastructure_external_nethttp.ClientHTTPInterface {
	if ClientHttpWithNet != nil {
		return ClientHttpWithNet
	} else {
		return &infrastructure_external_nethttp_nethttp.NetHTTPClient{}
	}
}

func InitContainer() {
	ClientHttpWithNet = ResolveClientHttpWithNet()
	DatabaseConnectionToLecWithPgx = ResolveDatabaseConnectionToLecWithPgx()
}
