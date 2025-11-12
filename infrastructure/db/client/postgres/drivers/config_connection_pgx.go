package infrastructure_db_client_postgres_drivers

import (
	"context"
	"fmt"
	"log"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
	"net/url"
	"regexp"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

const defaultMaxConns = int32(4)
const defaultMinConns = int32(0)
const defaultMaxConnLifetime = time.Hour
const defaultMaxConnIdleTime = time.Minute * 30
const defaultHealthCheckPeriod = time.Minute
const defaultConnectTimeout = time.Second * 5

type ConfigConnectionPgx struct {
	UrlToConnect string
	pool         *pgxpool.Pool
	Context      context.Context
}

func (configConnection *ConfigConnectionPgx) GetDatabaseConnection() (*infrastructure_db_client.ConfigDatabaseConnectionDriver, error) {

	err := configConnection.configuratePool()
	if err != nil {
		return nil, err
	}

	connection, err := configConnection.pool.Acquire(context.Background())

	if err != nil {
		return nil, err
	}

	return &infrastructure_db_client.ConfigDatabaseConnectionDriver{PgxConn: connection}, nil
}

func (configConnection *ConfigConnectionPgx) Select(ctx context.Context, dst interface{}, query string, args ...any) error {
	if err := configConnection.configuratePool(); err != nil {
		return err
	}

	connection, err := configConnection.pool.Acquire(context.Background())

	if err != nil {
		return err
	}

	defer connection.Release()

	if err := pgxscan.Select(ctx, connection, dst, query, args...); err != nil {
		return err
	}

	return nil
}

func CleanConectionString(rawConn string) (string, error) {
	re := regexp.MustCompile(`^postgres(?:ql)?://([^:@/]+):([^@]+)@([^/]+)(/[^?]+)?(\?.*)?$`)

	matches := re.FindStringSubmatch(rawConn)
	if len(matches) == 0 {
		return "", fmt.Errorf("formato inválido de conexión")
	}

	user := matches[1]
	pass := matches[2]
	hostPort := matches[3]
	dbPath := matches[4]
	query := matches[5]

	escapedUser := url.QueryEscape(user)
	escapedPass := url.QueryEscape(pass)

	finalURL := fmt.Sprintf("postgres://%s:%s@%s%s%s",
		escapedUser,
		escapedPass,
		hostPort,
		dbPath,
		query,
	)

	return finalURL, nil
}

func (configConnection *ConfigConnectionPgx) configuratePool() error {

	if configConnection.pool == nil {
		urlConnection, _err := CleanConectionString(configConnection.UrlToConnect)
		if _err != nil {
			return _err
		}

		if dbConfig, err := pgxpool.ParseConfig(urlConnection); err != nil {
			log.Println("Error al obtener conexion postgres")
			return err
		} else {
			dbConfig.MaxConns = defaultMaxConns
			dbConfig.MinConns = defaultMinConns
			dbConfig.MaxConnLifetime = defaultMaxConnLifetime
			dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
			dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
			dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout
			configConnection.pool, err = pgxpool.NewWithConfig(context.Background(), dbConfig)

			if err != nil {
				return err
			}
		}
	}

	return nil

}
