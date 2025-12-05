package repositories

import (
	"context"
	"encoding/json"

	"ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	value_constants "ms-genexis-pos-operaciones/context/creditcustomers/domain/value_object/constants"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
)

type PreAuthorizationRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *PreAuthorizationRepository) Create(ctx context.Context, req *entities.PreAuthorizationRequest) (bool, error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return false, err
	}
	defer conn.PgxConn.Release()

	payload, err := json.Marshal(req)
	if err != nil {
		return false, err
	}

	var created bool
	if err := conn.PgxConn.QueryRow(
		ctx,
		value_constants.QueryInsertPreAuthorization,
		string(payload),
	).Scan(&created); err != nil {
		return false, err
	}

	return created, nil
}
