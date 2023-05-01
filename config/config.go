package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// String de conexão com o banco de dados
	StringConnectionBD = ""

	// Porta onde a API vai estar rodando
	Port = 0

	//é a chave que será usada para criptografar o token
	SecretKey []byte
)

// init variables for the application
func LoadConfig() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Port, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Port = 9000
	}
	StringConnectionBD = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
	)
	fmt.Println("aqui")
	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
