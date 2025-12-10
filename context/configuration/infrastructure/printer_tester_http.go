package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
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

	url := os.Getenv("MICRO_PRINTER_URL")
	if url == "" {
		url = "http://localhost:18887/printer"
	}

	timeout := parseTimeout(os.Getenv("PRINTER_TIMEOUT_MS"))
	client := &http.Client{Timeout: timeout}

	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(payload))
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

func parseTimeout(msStr string) time.Duration {
	if msStr == "" {
		return 3 * time.Second
	}
	if ms, err := strconv.Atoi(msStr); err == nil && ms > 0 {
		return time.Duration(ms) * time.Millisecond
	}
	return 3 * time.Second
}
