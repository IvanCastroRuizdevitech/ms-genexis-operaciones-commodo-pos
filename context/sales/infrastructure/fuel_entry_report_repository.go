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

type FuelEntryReportRepository struct {
    Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *FuelEntryReportRepository) Generate(request *entities_sales.FuelEntryReportRequest) (*entities_main.Response[entities_sales.FuelEntryReportResult], error) {
    log.Println("CONSULTANDO:", constants_sales.QUERY_FUEL_ENTRY_REPORT)
    log.Println("ARGUMENTO 1 (numero_factura):", request.NumeroFactura)
    log.Println("ARGUMENTO 2 (copia):", request.Copia)
    log.Println("ARGUMENTO 3 (cola):", request.Cola)

    var rows []entities_sales.FuelEntryReportResult

    if err := r.Connection.Select(
        context.Background(),
        &rows,
        constants_sales.QUERY_FUEL_ENTRY_REPORT,
        request.NumeroFactura,
        request.Copia,
        request.Cola,
    ); err != nil {
        return nil, err
    }

    var row entities_sales.FuelEntryReportResult
    if len(rows) > 0 {
        row = rows[0]
    }

    success := entities_main.NewSuccessResponse[entities_sales.FuelEntryReportResult](
        200,
        "OK",
        time.Now().Format("2006-01-02 15:04:05"),
        &row,
    )

    return &success, nil
}

