package repositories

import (
	"context"
	"encoding/json"
	"log"
	"ms-genexis-pos-operaciones/context/envelopes/domain/entities"
	"ms-genexis-pos-operaciones/context/envelopes/domain/value_object/constants"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
)

type CreateEnvelopeRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (g *CreateEnvelopeRepository) Create(envelope_request *entities.EnvelopeRequest) (*entities.EnvelopeCreate, error) {

	conn, err := g.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}

	defer conn.PgxConn.Release()

	log.Println("CONSULTANDO: ", constants.QUERY_CREATE_ENVELOPE)
	jsonByte, _ := json.Marshal(envelope_request)
	log.Println("ARGUMENTO 1 : ", string(jsonByte))

	response := &entities.EnvelopeCreate{}
	err = conn.PgxConn.QueryRow(
		context.Background(),
		constants.QUERY_GET_TOTAL_ENVELOPES,
		envelope_request,
	).Scan(
		&response.Created,
		&response.MessageError,
	)

	if err != nil {
		return nil, err
	}
	jsonRespByte, _ := json.Marshal(response)
	log.Printf("Respuesta de la funcion: %s \n\n", jsonRespByte)

	return response, nil

}
