package main

import (
	"backend_api/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server is running in port 8000")
	r := routes.GenerateRouter()
	log.Fatal(http.ListenAndServe(":8000", r))
}
