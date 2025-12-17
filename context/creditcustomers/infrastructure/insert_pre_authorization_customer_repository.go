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

type InsertPreAuthorizationCustomerRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *InsertPreAuthorizationCustomerRepository) Insert(request *entities_creditcustomers.InsertPreAuthorizationCustomerRequest) (*entities_main.Response[entities_creditcustomers.InsertPreAuthorizationCustomerResult], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	payloadBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	var insertOk bool

	log.Println("CONSULTANDO:", constants_creditcustomers.QUERY_INSERT_PRE_AUTHORIZATION_CUSTOMER)
	log.Println("ARGUMENTO 1 (jsonb):", string(payloadBytes))

	err = conn.PgxConn.QueryRow(
		context.Background(),
		constants_creditcustomers.QUERY_INSERT_PRE_AUTHORIZATION_CUSTOMER,
		string(payloadBytes),
	).Scan(&insertOk)
	if err != nil {
		return nil, err
	}

	data := entities_creditcustomers.InsertPreAuthorizationCustomerResult{Success: insertOk}
	response := entities_main.NewSuccessResponse[entities_creditcustomers.InsertPreAuthorizationCustomerResult](
		200,
		"OK",
		time.Now().Format("2006-01-02 15:04:05"),
		&data,
	)

	return &response, nil
}
