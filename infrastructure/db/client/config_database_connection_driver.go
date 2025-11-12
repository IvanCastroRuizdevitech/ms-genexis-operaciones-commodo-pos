package infrastructure_db_client

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type ConfigDatabaseConnectionDriver struct {
	PgxConn *pgxpool.Conn
	//other drivers providers
}

func (configDatabaseDriver *ConfigDatabaseConnectionDriver) UsePgxDatabaseConnectionDriver() *pgxpool.Conn {
	return configDatabaseDriver.PgxConn
}
