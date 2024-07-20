package account

import (
	"auth/database"
	"errors"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type Manager struct {
	dataSource *database.Source
}

func NewAccountManager(dataSource *database.Source) *Manager {
	am := &Manager{dataSource}
	return am
}

func (am *Manager) AddAccount(firstName string, lastName string, username string, password string) bool {
	// Hash the password
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	// Create a database connection
	conn, _ := am.dataSource.Connect()
	defer conn.Close()

	stmt, _ := conn.Prepare("insert into users (first_name, last_name, username, password_hash) values($1, $2, $3, $4)")
	defer stmt.Close()

	_, err := stmt.Exec(firstName, lastName, username, passwordHash)
	return err == nil
}

func (am *Manager) GetAccount(username string) (Account, error) {
	conn, err := am.dataSource.Connect()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	stmt, _ := conn.Prepare(
		`select user_id, first_name, last_name, username, password_hash
				from users 
				where username = $1 and is_deleted=false`)

	defer stmt.Close()

	rows, _ := stmt.Query(username)
	defer rows.Close()

	if rows.Next() {
		var userId int
		var firstName, lastName, username, passwordHash string
		rows.Scan(&userId, &firstName, &lastName, &username, &passwordHash)
		ac := Account{userId, firstName, lastName, username, passwordHash}
		return ac, nil
	}

	return Account{}, errors.New("account not found")
}

func (am *Manager) AccountExists(username string) bool {
	_, err := am.GetAccount(username)
	return err == nil
}
