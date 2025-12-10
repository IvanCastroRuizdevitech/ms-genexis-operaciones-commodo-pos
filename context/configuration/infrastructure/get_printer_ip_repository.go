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

type GetPrinterIPRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *GetPrinterIPRepository) Get() (*entities_main.Response[entities.PrinterIPResult], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	var resultStr string

	log.Println("CONSULTANDO:", constants_configuration.QUERY_GET_PRINTER_IP)

	err = conn.PgxConn.QueryRow(
		context.Background(),
		constants_configuration.QUERY_GET_PRINTER_IP,
	).Scan(&resultStr)
	if err != nil {
		return nil, err
	}

	payload := entities.PrinterIPResult{}

	if resultStr != "" {
		tmp := map[string]any{}
		if err := json.Unmarshal([]byte(resultStr), &tmp); err == nil {
			if success, ok := tmp["success"].(bool); ok {
				payload.Success = success
			}
			if dataRaw, ok := tmp["data"].(map[string]any); ok {
				if val, ok := dataRaw["valor"].(string); ok {
					payload.Data.Valor = val
				}
				if codigo, ok := dataRaw["codigo"].(string); ok {
					payload.Data.Codigo = codigo
				}
				if idFloat, ok := dataRaw["id"].(float64); ok {
					payload.Data.ID = int(idFloat)
				}
			}
		} else {
			// if plain string, store as valor
			payload.Data.Valor = resultStr
		}
	}

	data := payload
	response := entities_main.NewSuccessResponse[entities.PrinterIPResult](
		200,
		"OK",
		time.Now().Format("2006-01-02 15:04:05"),
		&data,
	)

	return &response, nil
}
