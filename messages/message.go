package messages

import (
	"encoding/json"
	"log"
	"net/http"
)

// retorna uma resposta de mensagem
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		if erro := json.NewEncoder(w).Encode(data); erro != nil {
			log.Fatal(erro)
		}
	}
}

// retorna um erro de mensagem
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
