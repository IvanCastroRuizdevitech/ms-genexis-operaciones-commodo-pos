package api_adapter_server

import (
	"log"
	api_routes "ms-genexis-pos-operaciones/presentation/api/gin/routes"
	"os"
)

func Start() error {

	servidor, err := api_routes.GinConfig()

	if err != nil {
		log.Fatal(err)
		return err
	}

	port := ":" + os.Getenv("SERVER_PORT")
	dev := os.Getenv("DEV") == "true"

	if dev {
		log.Println("INICIANDO SERVIDOR SIN SSL en el puerto", port)
		if err := servidor.Run(port); err != nil {
			log.Fatal(err)
		}
	} else {
		certFile := os.Getenv("SSL_CERT_FILE")
		keyFile := os.Getenv("SSL_KEY_FILE")

		log.Println("INICIANDO SERVIDOR CON SSL en el puerto")
		if err := servidor.RunTLS(port, certFile, keyFile); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
