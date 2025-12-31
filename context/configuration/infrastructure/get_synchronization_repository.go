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

type GetSynchronizationRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *GetSynchronizationRepository) Get(idSincronizacion int, fechaInicio string, fechaFin string) (*entities_main.Response[[]entities.SynchronizationDetail], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	var jsonResult []byte

	log.Println("CONSULTANDO:", constants_configuration.QUERY_GET_SYNCHRONIZATION)
	log.Println("ARGUMENTOS:", idSincronizacion, fechaInicio, fechaFin)

	err = conn.PgxConn.QueryRow(
		context.Background(),
		constants_configuration.QUERY_GET_SYNCHRONIZATION,
		idSincronizacion,
		fechaInicio,
		fechaFin,
	).Scan(&jsonResult)
	if err != nil {
		return nil, err
	}

	if len(jsonResult) == 0 {
		jsonResult = []byte("[]")
	}

	var details []entities.SynchronizationDetail
	if err := json.Unmarshal(jsonResult, &details); err != nil {
		log.Println("[GetSynchronizationRepository][Unmarshal]", err)
		return nil, err
	}

	data := details
	success := entities_main.NewSuccessResponse[[]entities.SynchronizationDetail](
		200,
		"Sincronizacion obtenida correctamente",
		time.Now().Format("2006-01-02 15:04:05"),
		&data,
	)

	return &success, nil
}
