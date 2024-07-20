package page

import (
	"auth/account"
)

type UserData struct {
	LoggedIn bool
	Account  account.Account
}
