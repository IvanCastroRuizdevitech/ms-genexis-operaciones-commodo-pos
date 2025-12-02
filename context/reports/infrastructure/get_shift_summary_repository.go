package repositories

import (
	"context"
	"encoding/json"
	"log"
	"time"

	entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
	constants_reports "ms-genexis-pos-operaciones/context/reports/domain/value_object/constants"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
)

type GetShiftSummaryRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *GetShiftSummaryRepository) Get(pos int, fechaInicio string, fechaFin string) (*entities_main.Response[entities_reports.ShiftSummary], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	var jsonResult []byte

	log.Println("CONSULTANDO:", constants_reports.QUERY_GET_SHIFT_SUMMARY)
	log.Println("ARGUMENTOS:", pos, fechaInicio, fechaFin)

	err = conn.PgxConn.QueryRow(
		context.Background(),
		constants_reports.QUERY_GET_SHIFT_SUMMARY,
		pos,
		fechaInicio,
		fechaFin,
	).Scan(&jsonResult)
	if err != nil {
		return nil, err
	}

	if len(jsonResult) == 0 {
		jsonResult = []byte("[]")
	}

	var summary entities_reports.ShiftSummary
	if err := json.Unmarshal(jsonResult, &summary); err != nil {
		log.Println("[GetShiftSummaryRepository][Unmarshal]", err)
		return nil, err
	}

	data := summary
	success := entities_main.NewSuccessResponse[entities_reports.ShiftSummary](
		200,
		"Resumen obtenido correctamente",
		time.Now().Format("2006-01-02 15:04:05"),
		&data,
	)

	return &success, nil
}
