package repositories

import (
	"context"
	"encoding/json"
	"database/sql"
	"log"
	entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
	constants_reports "ms-genexis-pos-operaciones/context/reports/domain/value_object/constants"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
	"time"
)

type GetDayClosingReportRepository struct {
    Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *GetDayClosingReportRepository) GetReport(fecha string) (*entities_main.Response[entities_reports.DayClosingReport], error) {
    conn, err := r.Connection.GetDatabaseConnection()
    if err != nil {
        return nil, err
    }

    var jsonStrResponse sql.NullString
    defer conn.PgxConn.Release()

    log.Println("CONSULTANDO: ", constants_reports.QUERY_GET_DAY_CLOSING_REPORT)
    log.Println("ARGUMENTO 1 : ", fecha)

    err = conn.PgxConn.QueryRow(
        context.Background(),
        constants_reports.QUERY_GET_DAY_CLOSING_REPORT,
        fecha,
    ).Scan(&jsonStrResponse)

    if err != nil {
        return nil, err
    }

    // If DB returns NULL or "null", standardize to empty JSON object
    raw := "{}"
    if jsonStrResponse.Valid && jsonStrResponse.String != "" && jsonStrResponse.String != "null" {
        raw = jsonStrResponse.String
    }

    var data entities_reports.DayClosingReport
    if err := json.Unmarshal([]byte(raw), &data); err != nil {
        log.Println("[ERROR GetDayClosingReportRepository]", err)
        return nil, err
    }

    success := entities_main.NewSuccessResponse[entities_reports.DayClosingReport](
        200,
        "OK",
        time.Now().Format("2006-01-02 15:04:05"),
        &data,
    )

    log.Printf("Respuesta de la funcion despues del unmarshall: %+v \n\n", success)
    return &success, nil
}
