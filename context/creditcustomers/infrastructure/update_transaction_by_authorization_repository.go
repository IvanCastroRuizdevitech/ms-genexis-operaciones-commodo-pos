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

type UpdateTransactionByAuthorizationRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *UpdateTransactionByAuthorizationRepository) Update(request *entities_creditcustomers.UpdateTransactionByAuthorizationRequest) (*entities_main.Response[entities_creditcustomers.UpdateTransactionByAuthorizationResult], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	var jsonResult string

	tramaBytes, err := json.Marshal(request.Trama)
	if err != nil {
		return nil, err
	}

	log.Println("CONSULTANDO:", constants_creditcustomers.QUERY_UPDATE_TRANSACTION_BY_AUTHORIZATION)

	err = conn.PgxConn.QueryRow(
		context.Background(),
		constants_creditcustomers.QUERY_UPDATE_TRANSACTION_BY_AUTHORIZATION,
		request.Autorizacion,
		request.Surtidor,
		request.Cara,
		request.Grado,
		request.DocumentoCliente,
		request.PlacaVehiculo,
		request.MontoMaximo,
		request.CantidadMaxima,
		request.ClienteNombre,
		request.VehiculoOdometro,
		tramaBytes,
		int16(request.EstadoTransaccion),
		request.DocumentoConductor,
		request.ConductorNombre,
		request.ClienteTipoIdentificacionID,
	).Scan(&jsonResult)
	if err != nil {
		return nil, err
	}

	if jsonResult == "" {
		jsonResult = "{}"
	}

	payload := entities_creditcustomers.UpdateTransactionByAuthorizationResult{}
	if err := json.Unmarshal([]byte(jsonResult), &payload); err != nil {
		log.Println("[UpdateTransactionByAuthorizationRepository][Unmarshal]", err)
		return nil, err
	}

	data := payload
	response := entities_main.NewSuccessResponse[entities_creditcustomers.UpdateTransactionByAuthorizationResult](
		200,
		"OK",
		time.Now().Format("2006-01-02 15:04:05"),
		&data,
	)

	return &response, nil
}
