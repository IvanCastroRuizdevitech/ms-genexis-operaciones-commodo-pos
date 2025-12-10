package handler_configuration

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	container_configuration "ms-genexis-pos-operaciones/context/configuration/presentation/container"
	entities_main "ms-genexis-pos-operaciones/domain/entities"

	"github.com/gin-gonic/gin"
)

func UpdatePrinterIPHandler(ctx *gin.Context) {
	body := entities.PrinterIPUpdateRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, entities_main.NewErrorResponse[interface{}]("Payload invalido", err))
		return
	}

	response, err := container_configuration.ResolveUpdatePrinterIPContainer().Execute(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities_main.NewErrorResponse[interface{}]("Error interno", err))
		return
	}

	// Remap the response to the format required by the POS client.
	var parsed map[string]any
	if response.Data != nil && response.Data.Parsed != nil {
		parsed = response.Data.Parsed
	}

	dataPayload := map[string]any{}
	if parsed != nil {
		if dataRaw, ok := parsed["data"].(map[string]any); ok {
			if valorAnterior, ok := dataRaw["valor_anterior"]; ok {
				dataPayload["valor_anterior"] = valorAnterior
			}
			if valorNuevo, ok := dataRaw["valor_nuevo"]; ok {
				dataPayload["valor_nuevo"] = valorNuevo
			}
		}
	}

	mensaje := ""
	if parsed != nil {
		if msg, ok := parsed["mensaje"].(string); ok {
			mensaje = msg
		}
	}
	if mensaje == "" && response.Data != nil {
		mensaje = response.Data.Raw
	}

	hostDestino := ""
	if v, ok := dataPayload["valor_nuevo"].(string); ok && v != "" {
		hostDestino = v
	} else if v, ok := dataPayload["valor_anterior"].(string); ok && v != "" {
		hostDestino = v
	}

	if hostDestino != "" {
		if err := sendTestPrint(requestContext(ctx), hostDestino, 9100); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"success": false,
				"mensaje": "No se pudo imprimir la tirilla de prueba: " + err.Error(),
				"data":    dataPayload,
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  response.Status,
		"success": response.Success,
		"data":    dataPayload,
		"mensaje": mensaje,
	})
}

func sendTestPrint(ctx context.Context, host string, port int) error {
	type printRequest struct {
		Template []string       `json:"template"`
		Port     int            `json:"port"`
		Host     string         `json:"host"`
		Extra    map[string]any `json:"extraData,omitempty"`
	}

	tirilla := buildTestTemplate()
	body := printRequest{
		Template: tirilla,
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

	payload, err := jsonMarshal(body)
	if err != nil {
		return err
	}

	log.Printf("sendTestPrint start url=%s host=%s port=%d bytes=%d", url, host, port, len(payload))

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
		log.Printf("sendTestPrint error status=%d body=%s", resp.StatusCode, string(b))
		return fmt.Errorf("microservicio impresora devolvio %d: %s", resp.StatusCode, string(b))
	}

	log.Printf("sendTestPrint success status=%d", resp.StatusCode)
	return nil
}

func buildTestTemplate() []string {
	t := []string{}
	t = append(t, "\x1b\x61\x01")
	t = append(t, "************************************************\n\n")
	t = append(t, " CONFIGURACION CORRECTA\n\n")
	t = append(t, "************************************************\n\n")
	t = append(t, "\x1b\x61\x00")
	t = append(t, "\nEste es un test de impresion\n")
	t = append(t, "para verificar la conexion con\nla impresora Digital POS.\n")
	t = append(t, "Pangrama 1:\n")
	t = append(t, "El nino exclama de alegria viendo\nal fabuloso periquito comer\njugosos kiwis y zanahorias.\n")
	t = append(t, "Pangrama 2 (MAYUSCULAS):\n")
	t = append(t, "EL PINGUINO WENCESLAO HIZO KILOMETROS\nBAJO EXHAUSTIVA LLUVIA Y FRIO, ANORABA\nA SU QUERIDO CACHORRO.\n")
	t = append(t, "Otro ejemplo:\n")
	t = append(t, "Es extrano mojar queso en la cerveza\ny probar whisky de garrafa.\n")
	t = append(t, "Numeros:\n")
	t = append(t, "1234567890 0987654321\n")
	t = append(t, "89123749123401263418234091273\n")
	t = append(t, "Caracteres especiales:\n")
	t = append(t, "* = / + @ ( [ { } ] )\n")
	t = append(t, "\x1b\x61\x01")
	t = append(t, "************************************************\n\n")
	t = append(t, "  IMPRESION FINALIZADA\n\n")
	t = append(t, "************************************************\n\n")
	t = append(t, "\n\n\n\n")
	t = append(t, "\x1d\x56\x00")
	return t
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

func jsonMarshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func requestContext(ctx *gin.Context) context.Context {
	if ctx.Request != nil && ctx.Request.Context() != nil {
		return ctx.Request.Context()
	}
	return context.Background()
}
