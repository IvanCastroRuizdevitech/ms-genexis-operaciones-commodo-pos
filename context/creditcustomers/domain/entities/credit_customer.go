package entities

// CreditCustomer represents a customer with an enabled credit line.
type CreditCustomer struct {
	ID               int64   `json:"id"`
	Name             string  `json:"name"`
	Document         string  `json:"document"`
	CreditLimit      float64 `json:"credit_limit"`
	AvailableCredit  float64 `json:"available_credit"`
	Status           string  `json:"status"`
	LastUpdatedAtUtc string  `json:"last_updated_at_utc"`
}

// UpdateCreditCustomerLimitRequest contains the input required to adjust a customer's credit limit.
type UpdateCreditCustomerLimitRequest struct {
	CustomerID int64   `json:"customer_id" binding:"required"`
	NewLimit   float64 `json:"new_limit" binding:"required"`
}

// UpdateCreditCustomerLimitResult returns the updated limits so the caller can refresh its view.
type UpdateCreditCustomerLimitResult struct {
	CustomerID      int64   `json:"customer_id"`
	NewLimit        float64 `json:"new_limit"`
	AvailableCredit float64 `json:"available_credit"`
	Status          string  `json:"status"`
}
