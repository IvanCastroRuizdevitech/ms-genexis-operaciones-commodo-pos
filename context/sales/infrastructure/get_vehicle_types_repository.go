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

type GetVehicleTypesRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *GetVehicleTypesRepository) Get() (*entities_main.Response[entities_sales.VehicleTypesResponse], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	var jsonResult []byte

	log.Println("CONSULTANDO:", constants_sales.QUERY_GET_VEHICLE_TYPES)

	if err := conn.PgxConn.QueryRow(
		context.Background(),
		constants_sales.QUERY_GET_VEHICLE_TYPES,
	).Scan(&jsonResult); err != nil {
		return nil, err
	}

	if len(jsonResult) == 0 {
		jsonResult = []byte("{}")
	}

	result := entities_sales.VehicleTypesResponse{}
	if err := json.Unmarshal(jsonResult, &result); err != nil {
		log.Println("[GetVehicleTypesRepository][Unmarshal]", err)
		return nil, err
	}

	success := entities_main.NewSuccessResponse[entities_sales.VehicleTypesResponse](
		200,
		"OK",
		time.Now().Format("2006-01-02 15:04:05"),
		&result,
	)

	return &success, nil
}
