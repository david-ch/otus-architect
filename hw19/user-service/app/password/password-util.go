package password

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPwd calculates hash for the specified password
func HashPwd(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)

	if err != nil {
		log.Println("Error while calculating hash", err)
		return "", err
	}

	return string(hash), nil
}

// ValidatePwd compares password with password hash
func ValidatePwd(hash string, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}
