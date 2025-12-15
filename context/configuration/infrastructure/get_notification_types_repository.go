package repositories

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	constants_configuration "ms-genexis-pos-operaciones/context/configuration/domain/value_object/constants"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
)

type GetNotificationTypesRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *GetNotificationTypesRepository) GetNotificationTypes() (*entities_main.Response[entities.NotificationTypesResponse], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	jsonStrResponse := "{}"

	log.Println("CONSULTANDO:", constants_configuration.QUERY_GET_NOTIFICATION_TYPES)

	if err := conn.PgxConn.QueryRow(
		context.Background(),
		constants_configuration.QUERY_GET_NOTIFICATION_TYPES,
	).Scan(&jsonStrResponse); err != nil {
		return nil, err
	}

	payload := entities.NotificationTypesResponse{}
	if err := json.Unmarshal([]byte(jsonStrResponse), &payload); err != nil {
		log.Println("[ERROR GetNotificationTypesRepository]", err)
		return nil, err
	}

	success := entities_main.NewSuccessResponse[entities.NotificationTypesResponse](
		200,
		"OK",
		time.Now().Format("2006-01-02 15:04:05"),
		&payload,
	)

	return &success, nil
}
