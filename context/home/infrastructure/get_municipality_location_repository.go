package repositories

import (
	"context"
	"errors"
	"time"

	"ms-genexis-pos-operaciones/context/home/domain/entities"
	"ms-genexis-pos-operaciones/context/home/domain/value_object/constants"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"

	"github.com/jackc/pgx/v5"
)

type GetMunicipalityLocationRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *GetMunicipalityLocationRepository) GetByID(id int) (*entities_main.Response[entities.MunicipalityLocation], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	var location entities.MunicipalityLocation
	err = conn.PgxConn.QueryRow(
		context.Background(),
		constants.QUERY_GET_MUNICIPALITY_LOCATION,
		id,
	).Scan(
		&location.City,
		&location.Department,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			success := entities_main.NewSuccessResponse[entities.MunicipalityLocation](
				200,
				"No municipality found",
				time.Now().Format("2006-01-02 15:04:05"),
				nil,
			)
			return &success, nil
		}
		return nil, err
	}

	data := location
	success := entities_main.NewSuccessResponse[entities.MunicipalityLocation](
		200,
		"Municipality retrieved successfully",
		time.Now().Format("2006-01-02 15:04:05"),
		&data,
	)

	return &success, nil
}
