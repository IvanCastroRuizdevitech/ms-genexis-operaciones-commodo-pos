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

type GetInitialConfigurationRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *GetInitialConfigurationRepository) GetInitialConfiguration() (*entities_main.Response[entities.InitialConfiguration], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	jsonStr := "{}"
	log.Println("CONSULTANDO: ", constants.QUERY_GET_INITIAL_CONFIGURATION_POS)

	err = conn.PgxConn.QueryRow(
		context.Background(),
		constants.QUERY_GET_INITIAL_CONFIGURATION_POS,
	).Scan(&jsonStr)

	if err != nil {
		return nil, err
	}

	data := &entities.InitialConfiguration{}
	if err := json.Unmarshal([]byte(jsonStr), data); err != nil {
		log.Println("[ERROR GetInitialConfigurationRepository]", err)
		return nil, err
	}
	log.Printf("Resultado de la variable data: %+v\n", data)

	success := entities_main.NewSuccessResponse[entities.InitialConfiguration](
		200,
		"OK",
		time.Now().Format("2006-01-02 15:04:05"),
		data,
	)

	log.Printf("Respuesta de la funcion despues del unmarshall: %+v\n\n", success)

	return &success, nil
}
