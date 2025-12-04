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

type PriceFamiliesRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *PriceFamiliesRepository) GetAll(ctx context.Context) (*entities_main.Response[[]entities.PriceFamily], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	rows, err := conn.PgxConn.Query(ctx, value_constants.QueryGetPriceFamilies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []entities.PriceFamily
	for rows.Next() {
		m, err := pgx.RowToMap(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, entities.PriceFamily(m))
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
