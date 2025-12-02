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

type UpsertActiveDispenserFaceTransactionRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *UpsertActiveDispenserFaceTransactionRepository) Execute(surtidor int, cara int, codigo string, grado int, proveedorID int, montoMaximo float64, cantidadMaxima float64, trama map[string]any, promotorID int) (*entities_main.Response[map[string]any], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	var jsonResult []byte

	log.Println("CONSULTANDO:", constants_sales.QUERY_UPSERT_ACTIVE_DISPENSER_FACE_TRANSACTION)
	log.Println("ARGUMENTOS:", surtidor, cara, codigo, grado, proveedorID, montoMaximo, cantidadMaxima, promotorID)

	tramaBytes, err := json.Marshal(trama)
	if err != nil {
		return nil, err
	}

	err = conn.PgxConn.QueryRow(
		context.Background(),
		constants_sales.QUERY_UPSERT_ACTIVE_DISPENSER_FACE_TRANSACTION,
		surtidor,
		cara,
		codigo,
		grado,
		proveedorID,
		montoMaximo,
		cantidadMaxima,
		tramaBytes,
		promotorID,
	).Scan(&jsonResult)
	if err != nil {
		return nil, err
	}

	if len(jsonResult) == 0 {
		jsonResult = []byte("{}")
	}

	var transaction map[string]any
	if err := json.Unmarshal(jsonResult, &transaction); err != nil {
		log.Println("[UpsertActiveDispenserFaceTransactionRepository][Unmarshal]", err)
		return nil, err
	}

	data := transaction
	success := entities_main.NewSuccessResponse[map[string]any](
		200,
		"OK",
		time.Now().Format("2006-01-02 15:04:05"),
		&data,
	)

	return &success, nil
}
