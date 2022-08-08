package middleware

import (
	"backend_api/config"
	"backend_api/messages"
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}

}

// Verifica se o user fazendo a requisição está autenticado
func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := config.ValidateToken(r); err != nil {
			messages.Erro(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
