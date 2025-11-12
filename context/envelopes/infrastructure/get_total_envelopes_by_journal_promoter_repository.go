package repositories

import (
	"context"
	"encoding/json"
	"log"
	"ms-genexis-pos-operaciones/context/envelopes/domain/entities"
	"ms-genexis-pos-operaciones/context/envelopes/domain/value_object/constants"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
)

type GetTotalEnvelopesByJournalPromoterRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (g *GetTotalEnvelopesByJournalPromoterRepository) GetTotal(envelopes_request *entities.EnvelopesTotalRequest) (*entities_main.Response[entities.TotalEnvelopes], error) {

	conn, err := g.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}

	jsonStrResponse := "{}"

	defer conn.PgxConn.Release()

	log.Println("CONSULTANDO: ", constants.QUERY_GET_TOTAL_ENVELOPES)
	log.Println("ARGUMENTO 1 : ", string(envelopes_request.JournalId))
	log.Println("ARGUMENTO 2 : ", string(envelopes_request.PromoterId))

	err = conn.PgxConn.QueryRow(
		context.Background(),
		constants.QUERY_GET_TOTAL_ENVELOPES,
		envelopes_request.JournalId,
		envelopes_request.PromoterId,
	).Scan(&jsonStrResponse)

	if err != nil {
		return nil, err
	}
	log.Printf("Respuesta de la funcion: %s \n\n", jsonStrResponse)

	response := &entities_main.Response[entities.TotalEnvelopes]{}

	err = json.Unmarshal([]byte(jsonStrResponse), &response)

	if err != nil {
		log.Println("[ERROR ValidateClientQuotaRepository]", err)
		return nil, err
	}

	log.Printf("Respuesta de la funcion despues del unmarshall: %+v \n\n", response)

	return response, nil

}
