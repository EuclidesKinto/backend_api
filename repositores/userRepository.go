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
func (repository users) Create(user models.User) (uint64, error) {
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

// Busca todos os usuários
func (repository users) GetAll() ([]models.User, error) {
	var users []models.User
	rows, erro := repository.db.Query("SELECT id, name, email FROM users")
	if erro != nil {
		return nil, erro
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if erro = rows.Scan(&user.ID, &user.Name, &user.Email); erro != nil {
			return nil, erro
		}
		users = append(users, user)
	}
	return users, nil
}
