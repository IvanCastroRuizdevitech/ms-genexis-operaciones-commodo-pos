package usecase

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"ms-genexis-pos-operaciones/context/home/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/home/domain/ports/application/use_case"
	irepositories "ms-genexis-pos-operaciones/context/home/domain/ports/repositories"
	"ms-genexis-pos-operaciones/domain/constants"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	infrastructure_external_nethttp "ms-genexis-pos-operaciones/infrastructure/externals/externalhttp"
)

var _ iusecase.IProcessPendingTransmissions = (*ProcessPendingTransmissions)(nil)

type ProcessPendingTransmissions struct {
	FetchRepository  irepositories.IGetPendingTransmissionsRepository
	UpdateRepository irepositories.IUpdateTransmissionStatusRepository
	HTTPClient       infrastructure_external_nethttp.ClientHTTPInterface
}

func (u *ProcessPendingTransmissions) Execute() (*entities_main.Response[entities.TransmissionProcessSummary], error) {
	result, err := u.FetchRepository.GetAll()
	if err != nil {
		log.Println("[ProcessPendingTransmissions][Execute]", err)
		return nil, err
	}

	summary := entities.TransmissionProcessSummary{}
	if result == nil || result.Data == nil {
		success := entities_main.NewSuccessResponse(
			http.StatusOK,
			"No pending transmissions",
			time.Now().Format("2006-01-02 15:04:05"),
			&summary,
		)
		return &success, nil
	}

	for _, transmission := range *result.Data {
		summary.Processed++
		if err := u.processTransmission(transmission); err != nil {
			log.Printf("[ProcessPendingTransmissions] transmission %d failed: %v", transmission.IDTransmision, err)
			summary.Failed++
			continue
		}
		summary.Synchronized++
	}

	message := "Pending transmissions processed"
	if summary.Processed == 0 {
		message = "No pending transmissions"
	}

	success := entities_main.NewSuccessResponse(
		http.StatusOK,
		message,
		time.Now().Format("2006-01-02 15:04:05"),
		&summary,
	)

	return &success, nil
}

func (u *ProcessPendingTransmissions) processTransmission(transmission entities.PendingTransmission) error {
	method := strings.ToUpper(strings.TrimSpace(transmission.Method))
	var body []byte
	if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
		body = []byte(transmission.Request)
	}

	switch method {
	case http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch:
	default:
		return fmt.Errorf("unsupported method %s", method)
	}
	log.Printf("[ProcessPendingTransmissions] processing transmission %d to %s", transmission.IDTransmision, transmission.URL)
	log.Printf("[ProcessPendingTransmissions] request body: %s", string(body))
	headers := make(map[string]string, len(constants.DefaultHeaders))
	for k, v := range constants.DefaultHeaders {
		headers[k] = v
	}

	statusCode, responseBody, err := u.HTTPClient.Request(method, transmission.URL, body, headers)
	if err != nil {
		return err
	}

	log.Printf("[ProcessPendingTransmissions] transmission %d response status: %d", transmission.IDTransmision, statusCode)
	if len(responseBody) > 0 {
		log.Printf("[ProcessPendingTransmissions] transmission %d response body: %s", transmission.IDTransmision, string(responseBody))
	} else {
		log.Printf("[ProcessPendingTransmissions] transmission %d response body: <empty>", transmission.IDTransmision)
	}

	if statusCode < http.StatusOK || statusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("unexpected status code %d", statusCode)
	}

	return u.UpdateRepository.MarkSynchronized(transmission.IDTransmision)
}
