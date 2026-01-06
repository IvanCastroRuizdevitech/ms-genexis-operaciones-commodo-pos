package repositories

import (
	"context"
	"encoding/json"
	"log"
	"time"

	entities_comanda "ms-genexis-pos-operaciones/context/comanda/domain/entities"
	constants_comanda "ms-genexis-pos-operaciones/context/comanda/domain/value_object/constants"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
)

type UpdateComandaStatusRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *UpdateComandaStatusRepository) Update(request *entities_comanda.UpdateComandaStatusRequest) (*entities_main.Response[entities_comanda.UpdateComandaStatusResult], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	var jsonResult string

	log.Println("CONSULTANDO:", constants_comanda.QUERY_UPDATE_COMANDA_STATUS)

	err = conn.PgxConn.QueryRow(
		context.Background(),
		constants_comanda.QUERY_UPDATE_COMANDA_STATUS,
		request.ComandaID,
		request.NuevoEstadoID,
	).Scan(&jsonResult)
	if err != nil {
		return nil, err
	}

	if jsonResult == "" {
		jsonResult = "{}"
	}

	payload := entities_comanda.UpdateComandaStatusResult{}
	if err := json.Unmarshal([]byte(jsonResult), &payload); err != nil {
		log.Println("[UpdateComandaStatusRepository][Unmarshal]", err)
		return nil, err
	}

	response := entities_main.NewSuccessResponse[entities_comanda.UpdateComandaStatusResult](
		200,
		"OK",
		time.Now().Format("2006-01-02 15:04:05"),
		&payload,
	)

	return &response, nil
}
