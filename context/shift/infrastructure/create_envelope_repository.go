package repositories

import (
	"context"
	"encoding/json"
	"ms-genexis-pos-operaciones/context/shift/domain/entities"
	"ms-genexis-pos-operaciones/context/shift/domain/value_object/constants"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
)

type CreateEnvelopeRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *CreateEnvelopeRepository) Create(envelopeRequest *entities.EnvelopeRequest) (*entities.EnvelopeCreate, error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	payload, err := json.Marshal(envelopeRequest)
	if err != nil {
		return nil, err
	}

	response := &entities.EnvelopeCreate{}
	err = conn.PgxConn.QueryRow(
		context.Background(),
		constants.QUERY_CREATE_ENVELOPE,
		string(payload),
	).Scan(
		&response.Created,
		&response.MessageError,
	)
	if err != nil {
		return nil, err
	}

	if response.MessageError == "" {
		response.MessageError = "OK"
	}

	return response, nil
}
