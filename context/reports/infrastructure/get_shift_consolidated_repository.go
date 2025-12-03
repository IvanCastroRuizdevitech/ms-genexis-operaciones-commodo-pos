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

type GetShiftConsolidatedRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *GetShiftConsolidatedRepository) Get(fechaInicio string, fechaFin string) (*entities_main.Response[entities_reports.ShiftConsolidated], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	var jsonResult []byte

	log.Println("CONSULTANDO:", constants_reports.QUERY_GET_SHIFT_CONSOLIDATED)
	log.Println("ARGUMENTOS:", fechaInicio, fechaFin)

	err = conn.PgxConn.QueryRow(
		context.Background(),
		constants_reports.QUERY_GET_SHIFT_CONSOLIDATED,
		fechaInicio,
		fechaFin,
	).Scan(&jsonResult)
	if err != nil {
		return nil, err
	}

	if len(jsonResult) == 0 {
		jsonResult = []byte("[]")
	}

	var summary entities_reports.ShiftConsolidated
	if err := json.Unmarshal(jsonResult, &summary); err != nil {
		log.Println("[GetShiftConsolidatedRepository][Unmarshal]", err)
		return nil, err
	}

	data := summary
	success := entities_main.NewSuccessResponse[entities_reports.ShiftConsolidated](
		200,
		"Consolidado obtenido correctamente",
		time.Now().Format("2006-01-02 15:04:05"),
		&data,
	)

	return &success, nil
}
