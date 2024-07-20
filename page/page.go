package page

import (
	"auth/session"
)

type Data struct {
	Title string
	User  UserData
}

func NewData(title string, sm *session.Manager) *Data {
	userData := UserData{LoggedIn: sm.UserLoggedIn(), Account: sm.UserAccount()}
	return &Data{
		Title: title,
		User:  userData,
	}
}
