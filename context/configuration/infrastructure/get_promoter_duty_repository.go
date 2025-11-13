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

type GetPromoterDutyRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (g *GetPromoterDutyRepository) GetPromoterDuty() (*entities_main.Response[[]entities.PromoterDuty], error) {
	conn, err := g.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}

	jsonStrResponse := "[]"
	defer conn.PgxConn.Release()

	log.Println("CONSULTANDO: ", constants.QUERY_GET_PROMOTER_DUTY)

	err = conn.PgxConn.QueryRow(
		context.Background(),
		constants.QUERY_GET_PROMOTER_DUTY,
	).Scan(&jsonStrResponse)

	if err != nil {
		return nil, err
	}

	promoters := make([]entities.PromoterDuty, 0)

	err = json.Unmarshal([]byte(jsonStrResponse), &promoters)
	if err != nil {
		log.Println("[ERROR GetPromoterDutyRepository]", err)
		return nil, err
	}

	success := entities_main.NewSuccessResponse[[]entities.PromoterDuty](
		200,
		"OK",
		time.Now().Format("2006-01-02 15:04:05"),
		&promoters,
	)

	log.Printf("Respuesta de la funcion despues del unmarshall: %+v \n\n", success)

	return &success, nil
}
