package repositories

import (
	"context"
	"log"
	entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
	constants_reports "ms-genexis-pos-operaciones/context/reports/domain/value_object/constants"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
	"time"

	"github.com/jackc/pgx/v5"
)

type GetDailyNoveltiesRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *GetDailyNoveltiesRepository) GetNovelties(ano, mes, dia int) (*entities_main.Response[entities_reports.DailyNovelties], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	rows, err := conn.PgxConn.Query(
		context.Background(),
		constants_reports.QUERY_GET_DAILY_NOVELTIES,
		ano,
		mes,
		dia,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := pgx.CollectRows(rows, pgx.RowToMap)
	if err != nil {
		log.Println("[ERROR GetDailyNoveltiesRepository]", err)
		return nil, err
	}

	novelties := entities_reports.DailyNovelties(data)

	success := entities_main.NewSuccessResponse[entities_reports.DailyNovelties](
		200,
		"Novedades obtenidas correctamente",
		time.Now().Format("2006-01-02 15:04:05"),
		&novelties,
	)

	return &success, nil
}
