package usecase

import (
	"context"

	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	irepositories "ms-genexis-pos-operaciones/context/configuration/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

// PrintTestAfterUpdate takes the update response, remaps fields, triggers a test print, and returns the enriched response.
type PrintTestAfterUpdate struct {
	PrinterTester irepositories.IPrinterTester
}

func (u *PrintTestAfterUpdate) Execute(resp *entities_main.Response[entities.PrinterIPUpdateResult]) (*entities_main.Response[entities.PrinterIPUpdateResult], error) {
	if resp == nil || resp.Data == nil {
		return resp, nil
	}

	dataPayload, mensaje := remapPayload(resp.Data)
	resp.Data.Formatted = dataPayload
	resp.Data.Message = mensaje

	hostDestino := resolveHostDestino(dataPayload)
	if hostDestino != "" && u.PrinterTester != nil {
		if err := u.PrinterTester.PrintTest(context.Background(), hostDestino, 9100, buildTestTemplate()); err != nil {
			return nil, err
		}
	}

	return resp, nil
}

func remapPayload(result *entities.PrinterIPUpdateResult) (map[string]any, string) {
	dataPayload := map[string]any{}
	mensaje := ""

	var parsed map[string]any
	if result != nil && result.Parsed != nil {
		parsed = result.Parsed
	}

	if parsed != nil {
		if dataRaw, ok := parsed["data"].(map[string]any); ok {
			if valorAnterior, ok := dataRaw["valor_anterior"]; ok {
				dataPayload["valor_anterior"] = valorAnterior
			}
			if valorNuevo, ok := dataRaw["valor_nuevo"]; ok {
				dataPayload["valor_nuevo"] = valorNuevo
			}
		}
		if msg, ok := parsed["mensaje"].(string); ok {
			mensaje = msg
		}
	}

	if mensaje == "" && result != nil {
		mensaje = result.Raw
	}

	return dataPayload, mensaje
}

func resolveHostDestino(data map[string]any) string {
	if v, ok := data["valor_nuevo"].(string); ok && v != "" {
		return v
	}
	if v, ok := data["valor_anterior"].(string); ok && v != "" {
		return v
	}
	return ""
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
	t = append(t, "\n\n\n\n")
	t = append(t, "\n\n\n\n")
	t = append(t, "\n\n\n\n")
	t = append(t, "\x1d\x56\x00")
	return t
}
