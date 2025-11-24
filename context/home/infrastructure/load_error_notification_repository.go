package repositories

import (
	"context"
	"time"

	"ms-genexis-pos-operaciones/context/home/domain/entities"
	"ms-genexis-pos-operaciones/context/home/domain/value_object/constants"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
)

type LoadErrorNotificationRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *LoadErrorNotificationRepository) Load() (*entities_main.Response[[]entities.ErrorNotification], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	rows, err := conn.PgxConn.Query(context.Background(), constants.QUERY_LOAD_ERROR_NOTIFICATION)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notifications := make([]entities.ErrorNotification, 0, 16)
	for rows.Next() {
		var detail string
		if err := rows.Scan(&detail); err != nil {
			return nil, err
		}
		notifications = append(notifications, entities.ErrorNotification{Detail: detail})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	data := notifications
	success := entities_main.NewSuccessResponse[[]entities.ErrorNotification](
		200,
		"OK",
		time.Now().Format("2006-01-02 15:04:05"),
		&data,
	)

	return &success, nil
}
