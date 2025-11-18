package repositories

import (
    "context"
    "log"
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    constants_sales "ms-genexis-pos-operaciones/context/sales/domain/value_object/constants"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
    infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
    "time"
)

type UpdateClientMovementRepository struct {
    Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *UpdateClientMovementRepository) Update(request *entities_sales.UpdateClientMovementRequest) (*entities_main.Response[entities_sales.UpdateClientMovementResult], error) {
    log.Println("CONSULTANDO:", constants_sales.QUERY_UPDATE_CLIENT_MOVEMENT)
    log.Println("ARGUMENTO 1:", request.MovementId)
    log.Println("ARGUMENTO 2:", request.TransmissionId)
    log.Println("ARGUMENTO 3:", request.Synchronization)

    var rows []entities_sales.UpdateClientMovementResult

    if err := r.Connection.Select(
        context.Background(),
        &rows,
        constants_sales.QUERY_UPDATE_CLIENT_MOVEMENT,
        request.MovementId,
        request.TransmissionId,
        request.Synchronization,
    ); err != nil {
        return nil, err
    }

    var data entities_sales.UpdateClientMovementResult
    if len(rows) > 0 {
        data = rows[0]
    }

    success := entities_main.NewSuccessResponse[entities_sales.UpdateClientMovementResult](
        200,
        "OK",
        time.Now().Format("2006-01-02 15:04:05"),
        &data,
    )

    return &success, nil
}

