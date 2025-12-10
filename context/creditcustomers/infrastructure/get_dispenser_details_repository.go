package repositories

import (
	"context"
	"encoding/json"
	"log"
	"time"

	entities_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	constants_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/domain/value_object/constants"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
)

type GetDispenserDetailsRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *GetDispenserDetailsRepository) Get() (*entities_main.Response[entities_creditcustomers.DispenserDetailsFunctionResponse], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	var jsonResult string

	log.Println("CONSULTANDO:", constants_creditcustomers.QUERY_GET_DISPENSER_DETAILS)

	err = conn.PgxConn.QueryRow(
		context.Background(),
		constants_creditcustomers.QUERY_GET_DISPENSER_DETAILS,
	).Scan(&jsonResult)
	if err != nil {
		return nil, err
	}

	if jsonResult == "" {
		jsonResult = "{}"
	}

	payload := entities_creditcustomers.DispenserDetailsFunctionResponse{}
	if err := json.Unmarshal([]byte(jsonResult), &payload); err != nil {
		log.Println("[GetDispenserDetailsRepository][Unmarshal]", err)
		return nil, err
	}

	data := payload
	response := entities_main.NewSuccessResponse[entities_creditcustomers.DispenserDetailsFunctionResponse](
		200,
		"OK",
		time.Now().Format("2006-01-02 15:04:05"),
		&data,
	)

	return &response, nil
}
