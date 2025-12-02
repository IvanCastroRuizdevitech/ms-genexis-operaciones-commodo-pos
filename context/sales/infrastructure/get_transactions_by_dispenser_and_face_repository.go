package repositories

import (
	"context"
	"encoding/json"
	"log"
	"time"

	constants_sales "ms-genexis-pos-operaciones/context/sales/domain/value_object/constants"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
)

type GetTransactionsByDispenserAndFaceRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *GetTransactionsByDispenserAndFaceRepository) Get(dispenserID int, face int) (*entities_main.Response[[]map[string]interface{}], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	var jsonResult []byte

	log.Println("CONSULTANDO:", constants_sales.QUERY_GET_TRANSACTIONS_BY_DISPENSER_AND_FACE)
	log.Println("ARGUMENTO 1:", dispenserID)
	log.Println("ARGUMENTO 2:", face)

	err = conn.PgxConn.QueryRow(
		context.Background(),
		constants_sales.QUERY_GET_TRANSACTIONS_BY_DISPENSER_AND_FACE,
		dispenserID,
		face,
	).Scan(&jsonResult)
	if err != nil {
		return nil, err
	}

	if len(jsonResult) == 0 {
		jsonResult = []byte("[]")
	}

	var transactions []map[string]interface{}
	if err := json.Unmarshal(jsonResult, &transactions); err != nil {
		log.Println("[GetTransactionsByDispenserAndFaceRepository][Unmarshal]", err)
		return nil, err
	}

	data := transactions
	success := entities_main.NewSuccessResponse[[]map[string]interface{}](
		200,
		"OK",
		time.Now().Format("2006-01-02 15:04:05"),
		&data,
	)

	return &success, nil
}
