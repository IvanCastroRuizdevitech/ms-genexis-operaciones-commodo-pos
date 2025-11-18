package repositories

import (
	"context"
	"time"

	entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
	constants_reports "ms-genexis-pos-operaciones/context/reports/domain/value_object/constants"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"

	"github.com/jackc/pgx/v5"
)

type GetTanksRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *GetTanksRepository) GetAll() (*entities_main.Response[[]entities_reports.TankBodega], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	rows, err := conn.PgxConn.Query(
		context.Background(),
		constants_reports.QUERY_GET_TANKS,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := pgx.CollectRows(rows, pgx.RowToStructByName[entities_reports.TankBodega])
	if err != nil {
		return nil, err
	}

	success := entities_main.NewSuccessResponse[[]entities_reports.TankBodega](
		200,
		"Tanks successfully obtained",
		time.Now().Format("2006-01-02 15:04:05"),
		&data,
	)

	return &success, nil
}
