package repositories

import (
	"context"
	"encoding/json"
	"log"
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	"ms-genexis-pos-operaciones/context/configuration/domain/value_object/constants"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
	"time"
)

type ProcessNotificationRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *ProcessNotificationRepository) Process(request *entities.ProcessNotificationRequest) (*entities_main.Response[entities.ProcessNotificationResult], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}

	defer conn.PgxConn.Release()

	log.Println("CALL:", constants.QUERY_PROCESS_NOTIFICATION)
	log.Println("ARGUMENTO 1:", request.TipoNotificacion)
	log.Println("ARGUMENTO 2:", request.Data)
	log.Println("ARGUMENTO 3:", request.Prioridad)

	jsonResponse := "{}"

	if err := conn.PgxConn.QueryRow(
		context.Background(),
		constants.QUERY_PROCESS_NOTIFICATION,
		request.TipoNotificacion,
		request.Data,
		request.Prioridad,
		jsonResponse,
	).Scan(&jsonResponse); err != nil {
		return nil, err
	}

	var data entities.ProcessNotificationResult
	if err := json.Unmarshal([]byte(jsonResponse), &data); err != nil {
		return nil, err
	}

	success := entities_main.NewSuccessResponse[entities.ProcessNotificationResult](
		200,
		"OK",
		time.Now().Format("2006-01-02 15:04:05"),
		&data,
	)

	return &success, nil
}
