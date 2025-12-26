package repositories

import (
	"context"
	"encoding/json"
	"log"

	"ms-genexis-pos-operaciones/context/shift/domain/entities"
	"ms-genexis-pos-operaciones/context/shift/domain/value_object/constants"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
)

type GetPersonByIDRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *GetPersonByIDRepository) GetPersonByID(personaID int64) (*entities.PersonByIDResponse, error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	log.Println("CONSULTANDO:", constants.QUERY_GET_PERSON_BY_ID)
	log.Println("ARGUMENTO 1 : ", personaID)

	rawJSON := "{}"
	if err := conn.PgxConn.QueryRow(
		context.Background(),
		constants.QUERY_GET_PERSON_BY_ID,
		personaID,
	).Scan(&rawJSON); err != nil {
		return nil, err
	}

	result := &entities.PersonByIDResponse{}
	if err := json.Unmarshal([]byte(rawJSON), result); err != nil {
		return nil, err
	}

	return result, nil
}
