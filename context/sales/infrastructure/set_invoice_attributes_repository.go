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

type SetInvoiceAttributesRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *SetInvoiceAttributesRepository) Update(request *entities_sales.SetInvoiceAttributesRequest) (*entities_main.Response[entities_sales.SetInvoiceAttributesResult], error) {
	log.Println("CONSULTANDO:", constants_sales.QUERY_SET_INVOICE_ATTRIBUTES_BY_FACE)
	log.Println("ARGUMENTO 1 (cara):", request.Cara)

	var rows []entities_sales.SetInvoiceAttributesResult
	if err := r.Connection.Select(
		context.Background(),
		&rows,
		constants_sales.QUERY_SET_INVOICE_ATTRIBUTES_BY_FACE,
		request.Cara,
	); err != nil {
		return nil, err
	}

	var result entities_sales.SetInvoiceAttributesResult
	if len(rows) > 0 {
		result = rows[0]
	}

	success := entities_main.NewSuccessResponse[entities_sales.SetInvoiceAttributesResult](
		200,
		"OK",
		time.Now().Format("2006-01-02 15:04:05"),
		&result,
	)

	return &success, nil
}
