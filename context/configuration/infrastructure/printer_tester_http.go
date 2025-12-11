package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"ms-genexis-pos-operaciones/context/configuration/domain/value_object/constants"
	"net/http"
)

// PrinterTesterHTTP implements IPrinterTester using an HTTP call to the /printer microservice.
type PrinterTesterHTTP struct{}

type printerRequest struct {
	Template []string       `json:"template"`
	Port     int            `json:"port"`
	Host     string         `json:"host"`
	Extra    map[string]any `json:"extraData,omitempty"`
}

func (p *PrinterTesterHTTP) PrintTest(ctx context.Context, host string, port int, template []string) error {
	body := printerRequest{
		Template: template,
		Port:     port,
		Host:     host,
		Extra:    map[string]any{"imageBase64": ""},
	}

	client := &http.Client{}
	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, constants.MICRO_PRINTER_URL, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("microservicio impresora devolvio %d: %s", resp.StatusCode, string(b))
	}

	return nil
}
