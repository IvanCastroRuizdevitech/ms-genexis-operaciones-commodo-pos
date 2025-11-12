package repositories

import (
	"encoding/json"
	"log"
	"ms-genexis-pos-operaciones/context/shift/domain/entities"
	infrastructure_external_nethttp "ms-genexis-pos-operaciones/infrastructure/externals/externalhttp"
)

type SendOpeningShiftRepositoryHttp struct {
	Connection infrastructure_external_nethttp.ClientHTTPInterface
}

func (g *SendOpeningShiftRepositoryHttp) Send(url string, body []byte, header *map[string]string) (*entities.OpeningShiftHttpResponse, error) {
	log.Println("Send opening Shit URL => ", url)
	log.Println("Send opening Shit BODY => ", string(body))
	log.Println("Send opening Shit HEADERS => ", header)

	responseBody, err := g.Connection.Post(url, body, *header)
	if err != nil {
		return nil, err
	}

	var result entities.OpeningShiftHttpResponse
	log.Println("Response http => ", string(responseBody))

	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		log.Println("[ERROR Unmarshal Response]", err)
		return nil, err
	}

	return &result, nil
}
