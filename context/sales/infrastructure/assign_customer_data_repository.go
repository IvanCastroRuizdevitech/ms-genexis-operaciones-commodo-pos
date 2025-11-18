package repositories

import (
    "context"
    "encoding/json"
    "log"
    "strconv"
    "strings"
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    constants_sales "ms-genexis-pos-operaciones/context/sales/domain/value_object/constants"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
    infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
    "time"
)

type AssignCustomerDataRepository struct {
    Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *AssignCustomerDataRepository) Assign(request *entities_sales.AssignCustomerDataRequest) (*entities_main.Response[entities_sales.AssignCustomerDataResult], error) {
    payload, err := json.Marshal(request)
    if err != nil {
        return nil, err
    }

    log.Println("CONSULTANDO:", constants_sales.QUERY_ASSIGN_CUSTOMER_DATA)
    log.Println("ARGUMENTO 1 (json):", string(payload))

    var rows []entities_sales.AssignCustomerDataResult

    if err := r.Connection.Select(
        context.Background(),
        &rows,
        constants_sales.QUERY_ASSIGN_CUSTOMER_DATA,
        string(payload),
    ); err != nil {
        return nil, err
    }

    var result bool
    if len(rows) > 0 {
        switch v := rows[0].Info.(type) {
        case nil:
            result = false
        case bool:
            result = v
        case string:
            s := strings.TrimSpace(strings.ToLower(v))
            switch s {
            case "true", "t", "1", "yes", "y", "si":
                result = true
            case "false", "f", "0", "no", "n":
                result = false
            default:
                if i, err := strconv.ParseInt(s, 10, 64); err == nil {
                    result = i != 0
                } else if f, err := strconv.ParseFloat(s, 64); err == nil {
                    result = f != 0
                } else {
                    result = false
                }
            }
        case []byte:
            s := strings.TrimSpace(strings.ToLower(string(v)))
            switch s {
            case "true", "t", "1", "yes", "y", "si":
                result = true
            case "false", "f", "0", "no", "n":
                result = false
            default:
                if i, err := strconv.ParseInt(s, 10, 64); err == nil {
                    result = i != 0
                } else if f, err := strconv.ParseFloat(s, 64); err == nil {
                    result = f != 0
                } else {
                    result = false
                }
            }
        case int:
            result = v != 0
        case int8:
            result = v != 0
        case int16:
            result = v != 0
        case int32:
            result = v != 0
        case int64:
            result = v != 0
        case uint:
            result = v != 0
        case uint8:
            result = v != 0
        case uint16:
            result = v != 0
        case uint32:
            result = v != 0
        case uint64:
            result = v != 0
        case float32:
            result = v != 0
        case float64:
            result = v != 0
        default:
            result = false
        }
    }

    row := entities_sales.AssignCustomerDataResult{Info: result}

    success := entities_main.NewSuccessResponse[entities_sales.AssignCustomerDataResult](
        200,
        "OK",
        time.Now().Format("2006-01-02 15:04:05"),
        &row,
    )

    return &success, nil
}
