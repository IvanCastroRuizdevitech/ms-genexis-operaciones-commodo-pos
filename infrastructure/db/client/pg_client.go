package infrastructure_db_client

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ClientDb struct {
	Conn *pgxpool.Pool

	UrlConecion string
	contexto    context.Context
}

func (CDB *ClientDb) Select(query string, arguments []any) ([][]interface{}, error) {
	connecion, err := CDB.Conn.Acquire(CDB.contexto)

	if err != nil {
		return nil, err
	}
	defer connecion.Release()
	err = connecion.Ping(CDB.contexto)
	if err != nil {
		return nil, err
	}
	rows, err := connecion.Query(CDB.contexto, query, arguments...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var resultado [][]interface{}
	for rows.Next() {
		valores, err := rows.Values()

		if err != nil {
			return nil, err
		}
		resultado = append(resultado, valores)

	}

	return resultado, nil

}

func (CDB *ClientDb) Exec(query string, arguments []any) ([][]interface{}, error) {
	connecion, err := CDB.Conn.Acquire(CDB.contexto)
	if err != nil {
		return nil, fmt.Errorf("error al adquirir conexión: %w", err)
	}
	defer connecion.Release()

	if err := connecion.Ping(CDB.contexto); err != nil {
		return nil, fmt.Errorf("error al hacer ping a la base de datos: %w", err)
	}

	rows, err := connecion.Query(CDB.contexto, query, arguments...)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %w", err)
	}
	defer rows.Close()

	var resultado [][]interface{}
	for rows.Next() {
		valores, err := rows.Values()
		if err != nil {
			return nil, fmt.Errorf("error al leer valores de la fila: %w", err)
		}
		resultado = append(resultado, valores)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("error iterando sobre filas: %w", rows.Err())
	}

	return resultado, nil
}

const defaultMaxConns = int32(4)
const defaultMinConns = int32(0)
const defaultMaxConnLifetime = time.Hour
const defaultMaxConnIdleTime = time.Minute * 30
const defaultHealthCheckPeriod = time.Minute
const defaultConnectTimeout = time.Second * 5

func InitClient(UrlConn string) (*ClientDb, error) {
	cliente := &ClientDb{
		contexto:    context.Background(),
		UrlConecion: UrlConn,
	}
	if dbConfig, err := pgxpool.ParseConfig(cliente.UrlConecion); err != nil {
		return nil, err
	} else {
		dbConfig.MaxConns = defaultMaxConns
		dbConfig.MinConns = defaultMinConns
		dbConfig.MaxConnLifetime = defaultMaxConnLifetime
		dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
		dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
		dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout
		cliente.Conn, err = pgxpool.NewWithConfig(cliente.contexto, dbConfig)

		if err != nil {
			return nil, err
		}
	}

	return cliente, nil
}
