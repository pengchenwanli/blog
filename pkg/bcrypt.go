package pkg

import "golang.org/x/crypto/bcrypt"

func HashPassword(Password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
func ComparePassword(Password, HashPassword string) bool {
	err:=bcrypt.CompareHashAndPassword([]byte(Password),[]byte(HashPassword))
	return err==nil
}
