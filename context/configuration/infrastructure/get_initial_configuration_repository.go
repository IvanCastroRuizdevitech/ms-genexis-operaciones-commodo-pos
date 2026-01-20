package repositories

import (
	"context"
	"encoding/json"
	"log"
	"time"

	entities_configuration "ms-genexis-pos-operaciones/context/configuration/domain/entities"
	constants_configuration "ms-genexis-pos-operaciones/context/configuration/domain/value_object/constants"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
)

type GetInitialConfigurationRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *GetInitialConfigurationRepository) GetInitialConfiguration() (*entities_main.Response[entities_configuration.InitialConfigurationDataResponse], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	var jsonResult string

	log.Println("CONSULTANDO:", constants_configuration.QUERY_GET_ALL_CONFIGURATION)

	err = conn.PgxConn.QueryRow(
		context.Background(),
		constants_configuration.QUERY_GET_ALL_CONFIGURATION,
	).Scan(&jsonResult)
	if err != nil {
		return nil, err
	}

	if jsonResult == "" {
		jsonResult = "{}"
	}

	payload := entities_configuration.InitialConfigurationDataResponse{}
	if err := json.Unmarshal([]byte(jsonResult), &payload); err != nil {
		log.Println("[GetInitialConfigurationRepository][Unmarshal]", err)
		return nil, err
	}

	response := entities_main.NewSuccessResponse[entities_configuration.InitialConfigurationDataResponse](
		200,
		"OK",
		time.Now().Format("2006-01-02 15:04:05"),
		&payload,
	)

	return &response, nil
}
