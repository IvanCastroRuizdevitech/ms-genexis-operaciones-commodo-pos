package infrastructure_db_client

import (
	"context"
)

type DatabaseConnectionInterface interface {
	GetDatabaseConnection() (*ConfigDatabaseConnectionDriver, error)
	Select(ctx context.Context, dst interface{}, query string, args ...any) error
}
