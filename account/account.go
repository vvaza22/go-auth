package account

import "golang.org/x/crypto/bcrypt"

type Account struct {
	UserId       int
	FirstName    string
	LastName     string
	Username     string
	PasswordHash string
}

func (a *Account) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.PasswordHash), []byte(password))
	return err == nil
}
