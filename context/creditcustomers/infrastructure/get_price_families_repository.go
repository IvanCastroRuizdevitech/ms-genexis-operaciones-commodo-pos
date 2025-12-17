package repositories

import (
	"context"
	"errors"
	"log"
	"time"

	entities_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	constants_creditcustomers "ms-genexis-pos-operaciones/context/creditcustomers/domain/value_object/constants"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
)

type GetPriceFamiliesRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *GetPriceFamiliesRepository) Get() (*entities_main.Response[[]entities_creditcustomers.PriceFamily], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	log.Println("CONSULTANDO:", constants_creditcustomers.QUERY_GET_PRICE_FAMILIES)

	rows, err := conn.PgxConn.Query(context.Background(), constants_creditcustomers.QUERY_GET_PRICE_FAMILIES)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cols := rows.FieldDescriptions()
	if len(cols) == 0 {
		return nil, errors.New("sin columnas en el resultado")
	}

	results := make([]entities_creditcustomers.PriceFamily, 0, 16)
	for rows.Next() {
		vals, err := rows.Values()
		if err != nil {
			return nil, err
		}
		item := make(entities_creditcustomers.PriceFamily, len(cols))
		for i, col := range cols {
			item[string(col.Name)] = vals[i]
		}
		results = append(results, item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	data := results
	response := entities_main.NewSuccessResponse[[]entities_creditcustomers.PriceFamily](
		200,
		"OK",
		time.Now().Format("2006-01-02 15:04:05"),
		&data,
	)

	return &response, nil
}
