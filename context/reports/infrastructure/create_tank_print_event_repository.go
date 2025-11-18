package repositories

import (
	"context"
	"fmt"
	"log"
	"time"

	entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
	constants_reports "ms-genexis-pos-operaciones/context/reports/domain/value_object/constants"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
)

type CreateTankPrintEventRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *CreateTankPrintEventRepository) CreateEvent(tankIDs []int) (*entities_main.Response[entities_reports.TankPrintEventResult], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	if len(tankIDs) == 0 {
		return nil, fmt.Errorf("tank_ids array is empty")
	}

	log.Printf("CreateTankPrintEventRepository tank_ids: %v", tankIDs)

	_, err = conn.PgxConn.Exec(
		context.Background(),
		constants_reports.QUERY_CREATE_TANK_PRINT_EVENT,
		tankIDs,
	)
	if err != nil {
		return nil, err
	}

	result := entities_reports.TankPrintEventResult{Created: true}

	success := entities_main.NewSuccessResponse[entities_reports.TankPrintEventResult](
		200,
		"Tank print event successfully created.",
		time.Now().Format("2006-01-02 15:04:05"),
		&result,
	)

	return &success, nil
}
