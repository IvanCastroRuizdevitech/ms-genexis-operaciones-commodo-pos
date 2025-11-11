package api_adapter_server

import (
	"log"
	"ms-genexis-pos-operaciones/domain/constants"
	api_routes "ms-genexis-pos-operaciones/presentation/api/gin/routes"
	"os"
)

func Start() error {

	servidor, err := api_routes.GinConfig()
	if err != nil {
		log.Fatal(err)
		return err
	}

	port := ":" + os.Getenv(constants.HOST_PORT)

	log.Println("INICIANDO SERVIDOR SIN SSL en el puerto", port)
	if err := servidor.Run(port); err != nil {
		log.Fatal(err)
	}

	return nil
}
