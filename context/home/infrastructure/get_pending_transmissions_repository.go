package repositories

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"ms-genexis-pos-operaciones/context/home/domain/entities"
	"ms-genexis-pos-operaciones/context/home/domain/value_object/constants"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
)

type GetPendingTransmissionsRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *GetPendingTransmissionsRepository) GetAll() (*entities_main.Response[[]entities.PendingTransmission], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	var jsonResult []byte
	log.Println("CONSULTANDO:", constants.QUERY_GET_PENDING_TRANSMISSIONS)
	err = conn.PgxConn.QueryRow(
		context.Background(),
		constants.QUERY_GET_PENDING_TRANSMISSIONS,
	).Scan(&jsonResult)
	if err != nil {
		return nil, err
	}

	if len(jsonResult) == 0 {
		jsonResult = []byte("[]")
	}

	var transmissions []entities.PendingTransmission
	if err := json.Unmarshal(jsonResult, &transmissions); err != nil {
		log.Println("[GetPendingTransmissionsRepository][Unmarshal]", err)
		return nil, err
	}

	data := transmissions
	success := entities_main.NewSuccessResponse[[]entities.PendingTransmission](
		200,
		"Pending transmissions retrieved successfully",
		time.Now().Format("2006-01-02 15:04:05"),
		&data,
	)

	return &success, nil
}
