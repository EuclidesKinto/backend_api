package main

import (
	"backend_api/config"
	"backend_api/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()
	fmt.Println("Port:", config.Port)
	fmt.Println("StringConnectionBD:", config.StringConnectionBD)
	fmt.Println("Server is running in port 8000")
	r := routes.GenerateRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
