package repositores

import (
	"backend_api/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

// Cria um novo repositório de usuários
func NewRepositoryUsers(db *sql.DB) *users {
	return &users{db}
}

// Cria um novo usuário
func (repository *users) Create(user models.User) (uint64, error) {
	statement, erro := repository.db.Prepare(
		"INSERT INTO users (name, email, password) VALUES (?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Email, user.Password)
	if erro != nil {
		return 0, erro
	}
	lastIDInserted, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(lastIDInserted), nil

}
