package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		IfError(err)
	}

	return string(hashedPassword), nil

}

func HashToken(username string) string {

	hashToken, err := bcrypt.GenerateFromPassword([]byte(username), 1)
	IfError(err)

	return string(hashToken)

}
