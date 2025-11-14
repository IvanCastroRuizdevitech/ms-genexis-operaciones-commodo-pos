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

type CheckPendingSalesRepository struct {
    Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *CheckPendingSalesRepository) Check(request *entities_sales.CheckPendingSalesRequest) (*entities_main.Response[[]entities_sales.PendingSale], error) {
    log.Println("CONSULTANDO:", constants_sales.QUERY_CHECK_PENDING_SALES)
    log.Println("ARGUMENTO 1:", request.JournalId)
    log.Println("ARGUMENTO 2:", request.PromoterId)
    log.Println("ARGUMENTO 3:", request.Limit)

    var rows []entities_sales.PendingSale

    if err := r.Connection.Select(
        context.Background(),
        &rows,
        constants_sales.QUERY_CHECK_PENDING_SALES,
        request.JournalId,
        request.PromoterId,
        request.Limit,
    ); err != nil {
        return nil, err
    }

    success := entities_main.NewSuccessResponse[[]entities_sales.PendingSale](
        200,
        "OK",
        time.Now().Format("2006-01-02 15:04:05"),
        &rows,
    )

    return &success, nil
}

