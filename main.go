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
	r := routes.GenerateRouter()
	fmt.Println("Server is running in port :", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
