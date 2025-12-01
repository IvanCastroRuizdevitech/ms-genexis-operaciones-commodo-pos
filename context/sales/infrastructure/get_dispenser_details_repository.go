package repositories

import (
	"context"
	"encoding/json"
	"log"
	"time"

	entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
	constants_sales "ms-genexis-pos-operaciones/context/sales/domain/value_object/constants"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
)

type GetDispenserDetailsRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *GetDispenserDetailsRepository) Get() (*entities_main.Response[[]entities_sales.DispenserDetail], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	var jsonResult []byte

	log.Println("CONSULTANDO:", constants_sales.QUERY_GET_DISPENSER_DETAILS)

	err = conn.PgxConn.QueryRow(
		context.Background(),
		constants_sales.QUERY_GET_DISPENSER_DETAILS,
	).Scan(&jsonResult)
	if err != nil {
		return nil, err
	}

	if len(jsonResult) == 0 {
		jsonResult = []byte("[]")
	}

	var details []entities_sales.DispenserDetail
	if err := json.Unmarshal(jsonResult, &details); err != nil {
		log.Println("[GetDispenserDetailsRepository][Unmarshal]", err)
		return nil, err
	}

	data := details
	success := entities_main.NewSuccessResponse[[]entities_sales.DispenserDetail](
		200,
		"OK",
		time.Now().Format("2006-01-02 15:04:05"),
		&data,
	)

	return &success, nil
}
