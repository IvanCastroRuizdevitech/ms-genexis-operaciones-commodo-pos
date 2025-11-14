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

type CheckDatafonoCancellationsInProgressRepository struct {
    Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *CheckDatafonoCancellationsInProgressRepository) Check(request *entities_sales.DatafonoCancellationsInProgressRequest) (*entities_main.Response[entities_sales.DatafonoCancellationsInProgress], error) {
    log.Println("CONSULTANDO:", constants_sales.QUERY_CHECK_DATAFONO_CANCELLATIONS_IN_PROGRESS)
    log.Println("ARGUMENTO 1:", request.MovementId)
    log.Println("ARGUMENTO 2:", request.TransactionOperationId)
    log.Println("ARGUMENTO 3:", request.TransactionStatusId)

    var rows []entities_sales.DatafonoCancellationsInProgress

    if err := r.Connection.Select(
        context.Background(),
        &rows,
        constants_sales.QUERY_CHECK_DATAFONO_CANCELLATIONS_IN_PROGRESS,
        request.MovementId,
        request.TransactionOperationId,
        request.TransactionStatusId,
    ); err != nil {
        return nil, err
    }

    var row entities_sales.DatafonoCancellationsInProgress
    if len(rows) > 0 {
        row = rows[0]
    }

    success := entities_main.NewSuccessResponse[entities_sales.DatafonoCancellationsInProgress](
        200,
        "OK",
        time.Now().Format("2006-01-02 15:04:05"),
        &row,
    )

    return &success, nil
}

