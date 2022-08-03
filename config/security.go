package config

import "golang.org/x/crypto/bcrypt"

// hashPassword recebe uma string e coloca um hash na mesma
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// comparePassword recebe uma string e compara com um hash
func VerifyPassword(passwordHash, passwordString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(passwordString))
}
