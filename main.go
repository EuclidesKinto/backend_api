package main

import (
	"backend_api/config"
	"backend_api/routes"
	"fmt"
	"log"
	"net/http"
)

// func init() {
// 	key := make([]byte, 64)
// 	if _, err := rand.Read(key); err != nil {
// 		log.Fatal(err)
// 	}
// 	stringBase64 := base64.StdEncoding.EncodeToString(key)
// 	fmt.Println(stringBase64)
// }

func main() {
	config.LoadConfig()
	fmt.Println("Port:", config.Port)
	r := routes.GenerateRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
