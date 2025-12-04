package repositories

import (
	"context"
	"time"

	"ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	value_constants "ms-genexis-pos-operaciones/context/creditcustomers/domain/value_object/constants"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"

	"github.com/jackc/pgx/v5"
)

type IdentifierTypesRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *IdentifierTypesRepository) GetAll(ctx context.Context) (*entities_main.Response[[]entities.IdentifierType], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	rows, err := conn.PgxConn.Query(ctx, value_constants.QueryGetIdentifierTypes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []entities.IdentifierType
	for rows.Next() {
		m, err := pgx.RowToMap(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, entities.IdentifierType(m))
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	success := entities_main.NewSuccessResponse(
		200,
		"OK",
		time.Now().Format("2006-01-02 15:04:05"),
		&result,
	)

	return &success, nil
}
