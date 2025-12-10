package repositories

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	constants_configuration "ms-genexis-pos-operaciones/context/configuration/domain/value_object/constants"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
)

type UpdatePrinterIPRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *UpdatePrinterIPRepository) Update(req *entities.PrinterIPUpdateRequest) (*entities_main.Response[entities.PrinterIPUpdateResult], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	var resultStr string

	log.Println("CONSULTANDO:", constants_configuration.QUERY_UPDATE_PRINTER_IP)
	log.Println("ARGUMENTO IP:", req.IP)

	err = conn.PgxConn.QueryRow(
		context.Background(),
		constants_configuration.QUERY_UPDATE_PRINTER_IP,
		req.IP,
	).Scan(&resultStr)
	if err != nil {
		return nil, err
	}

	payload := entities.PrinterIPUpdateResult{}
	parsed := map[string]any{}
	if resultStr != "" {
		if err := json.Unmarshal([]byte(resultStr), &parsed); err == nil {
			payload.Parsed = parsed
			if msg, ok := parsed["mensaje"].(string); ok {
				payload.Raw = msg
			}
			if dataRaw, ok := parsed["data"].(map[string]any); ok {
				if val, ok := dataRaw["valor_nuevo"].(string); ok && val != "" {
					payload.Raw = val
				}
			}
		} else {
			payload.Raw = resultStr
		}
	} else {
		payload.Raw = resultStr
	}

	data := payload
	response := entities_main.NewSuccessResponse[entities.PrinterIPUpdateResult](
		200,
		"OK",
		time.Now().Format("2006-01-02 15:04:05"),
		&data,
	)

	return &response, nil
}
