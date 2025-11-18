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

type GetPendingSaleDatafonoRepository struct {
    Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *GetPendingSaleDatafonoRepository) Get(transactionId int) (*entities_main.Response[entities_sales.PendingSaleDatafono], error) {
    log.Println("CONSULTANDO:", constants_sales.QUERY_GET_PENDING_SALE_DATAFONO)
    log.Println("ARGUMENTO 1:", transactionId)

    var rows []entities_sales.PendingSaleDatafono

    if err := r.Connection.Select(
        context.Background(),
        &rows,
        constants_sales.QUERY_GET_PENDING_SALE_DATAFONO,
        transactionId,
    ); err != nil {
        return nil, err
    }

    var row entities_sales.PendingSaleDatafono
    if len(rows) > 0 {
        row = rows[0]
    }

    success := entities_main.NewSuccessResponse[entities_sales.PendingSaleDatafono](
        200,
        "OK",
        time.Now().Format("2006-01-02 15:04:05"),
        &row,
    )

    return &success, nil
}

