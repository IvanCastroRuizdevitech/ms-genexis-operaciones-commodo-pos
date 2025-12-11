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

type GetUnifiedConsecutivesRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *GetUnifiedConsecutivesRepository) GetUnifiedConsecutives() (*entities_main.Response[entities.UnifiedConsecutivesResponse], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	jsonStrResponse := "{}"

	log.Println("CONSULTANDO:", constants_configuration.QUERY_GET_UNIFIED_CONSECUTIVES)

	err = conn.PgxConn.QueryRow(
		context.Background(),
		constants_configuration.QUERY_GET_UNIFIED_CONSECUTIVES,
	).Scan(&jsonStrResponse)

	if err != nil {
		return nil, err
	}

	payload := entities.UnifiedConsecutivesResponse{}
	if err := json.Unmarshal([]byte(jsonStrResponse), &payload); err != nil {
		log.Println("[ERROR GetUnifiedConsecutivesRepository]", err)
		return nil, err
	}

	success := entities_main.NewSuccessResponse[entities.UnifiedConsecutivesResponse](
		200,
		"OK",
		time.Now().Format("2006-01-02 15:04:05"),
		&payload,
	)

	return &success, nil
}
