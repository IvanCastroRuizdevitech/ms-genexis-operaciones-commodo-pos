package repositories

import (
    "context"
    "errors"
    "log"
    "time"
    "ms-genexis-pos-operaciones/context/shift/domain/value_object/constants"
    infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
)

type GetDailyIncomeMeasurementsRepository struct {
    Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *GetDailyIncomeMeasurementsRepository) GetDailyIncomeMeasurements(start time.Time, end time.Time) ([]map[string]interface{}, error) {
    conn, err := r.Connection.GetDatabaseConnection()
    if err != nil {
        return nil, err
    }
    defer conn.PgxConn.Release()

    log.Println("CONSULTANDO: ", constants.QUERY_GET_DAILY_INCOME_MEASUREMENTS)
    log.Println("ARGUMENTO 1 : ", start)
    log.Println("ARGUMENTO 2 : ", end)

    rows, err := conn.PgxConn.Query(context.Background(), constants.QUERY_GET_DAILY_INCOME_MEASUREMENTS, start, end)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    cols := rows.FieldDescriptions()
    if len(cols) == 0 {
        return nil, errors.New("sin columnas en el resultado")
    }

    results := make([]map[string]interface{}, 0, 32)
    for rows.Next() {
        vals, err := rows.Values()
        if err != nil {
            return nil, err
        }
        m := make(map[string]interface{}, len(cols))
        for i, c := range cols {
            m[string(c.Name)] = vals[i]
        }
        results = append(results, m)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }

    return results, nil
}

