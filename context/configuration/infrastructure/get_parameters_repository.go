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

type GetParametersRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (g *GetParametersRepository) GetParameters() (*entities_main.Response[entities.Config], error) {

	conn, err := g.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}

	jsonStrResponse := "{}"

	defer conn.PgxConn.Release()

	log.Println("CONSULTANDO: ", constants.QUERY_GET_ALL_CONFIGURATION)

	err = conn.PgxConn.QueryRow(
		context.Background(),
		constants.QUERY_GET_ALL_CONFIGURATION,
	).Scan(&jsonStrResponse)

	if err != nil {
		return nil, err
	}

	config := &entities.Config{}

	err = json.Unmarshal([]byte(jsonStrResponse), config)

	if err != nil {
		log.Println("[ERROR ValidateClientQuotaRepository]", err)
		return nil, err
	}

	success := entities_main.NewSuccessResponse[entities.Config](
		200,
		"OK",
		time.Now().Format("2006-01-02 15:04:05"),
		config,
	)

	log.Printf("Respuesta de la funcion despues del unmarshall: %+v \n\n", success)

	return &success, nil

}
