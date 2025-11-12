package main

import (
	"log"
	"ms-genexis-pos-operaciones/domain/constants"
	"ms-genexis-pos-operaciones/presentation"
	api_adapter_server "ms-genexis-pos-operaciones/presentation/api/gin/adapter"
	presentation_container "ms-genexis-pos-operaciones/presentation/container"
	"os"
)

func main() {

	log.Println(constants.Green+"[¡MS GENEXIS POS OPERACIONES!] - Initialize server...", constants.Reset)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetOutput(os.Stdout)

	presentation_container.InitContainer()

	if err := api_adapter_server.Start(); err != nil {
		log.Fatal(err)
		panic(err)
	}

	presentation.Run()

}
