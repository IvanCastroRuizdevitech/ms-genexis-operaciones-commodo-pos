package repositories

import (
	"context"
	"encoding/json"
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	"ms-genexis-pos-operaciones/context/configuration/domain/value_object/constants"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
	"time"
)

type ValidateAdminPersonRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *ValidateAdminPersonRepository) ValidateAdminPerson(request *entities.AdminValidationRequest) (*entities_main.Response[entities.AdminValidationResult], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	rawJSON := "{}"

	if err := conn.PgxConn.QueryRow(
		context.Background(),
		constants.QUERY_VALIDATE_ADMIN_PERSON,
		request.Usuario,
		request.Clave,
		request.Tag,
	).Scan(&rawJSON); err != nil {
		return nil, err
	}

	result := entities.AdminValidationResult{}
	if err := json.Unmarshal([]byte(rawJSON), &result); err != nil {
		return nil, err
	}

	response := entities_main.Response[entities.AdminValidationResult]{
		Status:      200,
		Success:     result.Success,
		Message:     result.Message,
		ProcessDate: time.Now().Format("2006-01-02 15:04:05"),
		Data:        &result,
	}

	return &response, nil
}
