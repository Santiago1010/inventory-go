package encryption

import "golang.org/x/crypto/bcrypt"

const key = "4003A79E01752D9C0FDC00F0C203227B"

func HashPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}

func ValidatePassword(password string, hashedPassword []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		return err
	}
	return nil
}
