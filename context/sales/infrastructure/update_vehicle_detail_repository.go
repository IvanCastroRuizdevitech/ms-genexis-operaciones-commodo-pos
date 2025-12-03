package repositories

import (
	"context"
	"log"
	"time"

	entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
	constants_sales "ms-genexis-pos-operaciones/context/sales/domain/value_object/constants"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
)

type UpdateVehicleDetailRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *UpdateVehicleDetailRepository) Update(request *entities_sales.UpdateVehicleDetailRequest) (*entities_main.Response[entities_sales.UpdateVehicleDetailResult], error) {
	log.Println("CONSULTANDO:", constants_sales.QUERY_UPDATE_VEHICLE_DETAIL)
	log.Println("ARGUMENTO 1:", request.MovimientoID)
	log.Println("ARGUMENTO 2:", request.VehiculoPlaca)
	log.Println("ARGUMENTO 3:", request.VehiculoNumero)
	log.Println("ARGUMENTO 4:", request.VehiculoOdometro)

	var rows []entities_sales.UpdateVehicleDetailResult
	if err := r.Connection.Select(
		context.Background(),
		&rows,
		constants_sales.QUERY_UPDATE_VEHICLE_DETAIL,
		request.MovimientoID,
		request.VehiculoPlaca,
		request.VehiculoNumero,
		request.VehiculoOdometro,
	); err != nil {
		return nil, err
	}

	var result entities_sales.UpdateVehicleDetailResult
	if len(rows) > 0 {
		result = rows[0]
	}

	success := entities_main.NewSuccessResponse[entities_sales.UpdateVehicleDetailResult](
		200,
		"OK",
		time.Now().Format("2006-01-02 15:04:05"),
		&result,
	)

	return &success, nil
}
