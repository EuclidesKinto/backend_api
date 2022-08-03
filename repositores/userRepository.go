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

// Busca um usuário pelo ID
func (repository users) GetByID(ID uint64) (models.User, error) {
	var user models.User
	rows, erro := repository.db.Query("SELECT id, name, email FROM users WHERE id = ?", ID)
	if erro != nil {
		return models.User{}, erro
	}
	defer rows.Close()

	for rows.Next() {
		if erro = rows.Scan(&user.ID, &user.Name, &user.Email); erro != nil {
			return models.User{}, erro
		}
	}
	return user, nil
}

// Update um usuário
func (repository users) Update(ID uint64, user models.User) error {
	statement, erro := repository.db.Prepare(
		"UPDATE users SET name = ?, email = ? WHERE id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(user.Name, user.Email, ID); erro != nil {
		return erro
	}
	return nil
}

// Delete um usuário
func (repository users) Delete(ID uint64) error {
	statement, erro := repository.db.Prepare("DELETE FROM users WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}
	return nil
}

// Busca um usuário pelo email
func (repository users) GetByEmail(email string) (models.User, error) {
	var user models.User
	rows, erro := repository.db.Query("SELECT id, password FROM users WHERE email = ?", email)
	if erro != nil {
		return models.User{}, erro
	}
	defer rows.Close()

	for rows.Next() {
		if erro = rows.Scan(&user.ID, &user.Password); erro != nil {
			return models.User{}, erro
		}
	}
	return user, nil
}
