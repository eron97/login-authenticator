package crypto

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPassword(plainPassword, hashedPassword string) (bool, error) {
	hashedPasswordBytes := []byte(hashedPassword)

	if err := bcrypt.CompareHashAndPassword(hashedPasswordBytes, []byte(plainPassword)); err != nil {
		return false, err
	}

	return true, nil
}
