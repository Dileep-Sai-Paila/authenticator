package utils

import "golang.org/x/crypto/bcrypt"

// it takes a plaintext password and returns its bcrypt hash
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // takes the password as a byte slice and a cost factor
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// compares a plaintext password with a bcrypt hash to see if they match
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil // Returns true if err is nil (match), false otherwise
}
