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

type UpdateMovementStateRepository struct {
    Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *UpdateMovementStateRepository) Update(movementId int, request *entities_sales.UpdateMovementStateRequest) (*entities_main.Response[entities_sales.UpdateMovementStateResult], error) {
    conn, err := r.Connection.GetDatabaseConnection()
    if err != nil {
        return nil, err
    }
    defer conn.PgxConn.Release()

    log.Println("CONSULTANDO:", constants_sales.QUERY_UPDATE_MOVEMENT_STATE)
    log.Println("ARGUMENTO 1:", request.EstadoDianId)
    log.Println("ARGUMENTO 2:", request.EstadoDianId)
    log.Println("ARGUMENTO 3:", movementId)

    cmdTag, err := conn.PgxConn.Exec(
        context.Background(),
        constants_sales.QUERY_UPDATE_MOVEMENT_STATE,
        request.EstadoDianId,
        request.EstadoDianId,
        movementId,
    )
    if err != nil {
        return nil, err
    }

    updated := cmdTag.RowsAffected() > 0
    data := entities_sales.UpdateMovementStateResult{Updated: updated}

    success := entities_main.NewSuccessResponse[entities_sales.UpdateMovementStateResult](
        200,
        "OK",
        time.Now().Format("2006-01-02 15:04:05"),
        &data,
    )

    return &success, nil
}

