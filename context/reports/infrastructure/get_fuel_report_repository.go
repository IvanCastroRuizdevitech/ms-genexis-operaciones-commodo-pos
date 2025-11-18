package repositories

import (
	"context"
	entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
	constants_reports "ms-genexis-pos-operaciones/context/reports/domain/value_object/constants"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
	"time"

	"github.com/jackc/pgx/v5"
)

type GetFuelReportRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *GetFuelReportRepository) GetReport(fecha string) (*entities_main.Response[entities_reports.FuelReport], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	var report entities_reports.FuelReport
	err = conn.PgxConn.QueryRow(
		context.Background(),
		constants_reports.QUERY_GET_FUEL_REPORT,
		fecha,
	).Scan(
		&report.TotalVentasCombustible,
		&report.TotalVentasCanastilla,
		&report.CantidadVentasCombustible,
		&report.CantidadVentasCanastilla,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			report = entities_reports.FuelReport{}
		} else {
			return nil, err
		}
	}

	success := entities_main.NewSuccessResponse[entities_reports.FuelReport](
		200,
		"OK",
		time.Now().Format("2006-01-02 15:04:05"),
		&report,
	)

	return &success, nil
}
