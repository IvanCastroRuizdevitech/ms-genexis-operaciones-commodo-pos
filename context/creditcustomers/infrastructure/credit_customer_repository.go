package repositories

import (
	"context"
	"time"

	"ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	value_constants "ms-genexis-pos-operaciones/context/creditcustomers/domain/value_object/constants"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
)

// CreditCustomerRepository wires credit customer queries. Currently serves stub data until a real source is connected.
type CreditCustomerRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func nowString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (r *CreditCustomerRepository) GetAll(ctx context.Context) (*entities_main.Response[[]entities.CreditCustomer], error) {
	_ = ctx // reserved for DB usage

	customers := []entities.CreditCustomer{
		{
			ID:               1,
			Name:             "Default Credit Customer",
			Document:         "N/A",
			CreditLimit:      1000,
			AvailableCredit:  1000,
			Status:           value_constants.StatusActive,
			LastUpdatedAtUtc: time.Now().UTC().Format(time.RFC3339),
		},
	}

	success := entities_main.NewSuccessResponse[[]entities.CreditCustomer](
		200,
		"Credit customers fetched",
		nowString(),
		&customers,
	)

	return &success, nil
}

func (r *CreditCustomerRepository) UpdateLimit(ctx context.Context, request *entities.UpdateCreditCustomerLimitRequest) (*entities_main.Response[entities.UpdateCreditCustomerLimitResult], error) {
	_ = ctx // reserved for DB usage

	result := entities.UpdateCreditCustomerLimitResult{
		CustomerID:      request.CustomerID,
		NewLimit:        request.NewLimit,
		AvailableCredit: request.NewLimit,
		Status:          value_constants.StatusActive,
	}

	success := entities_main.NewSuccessResponse(
		200,
		"Credit limit updated (stub)",
		nowString(),
		&result,
	)

	return &success, nil
}
