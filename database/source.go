package database

import "database/sql"

type Source struct {
	host     string
	username string
	password string
	database string
}

func NewSource(host string, username string, password string, database string) *Source {
	return &Source{host: host, username: username, password: password, database: database}
}

func (dbs *Source) Connect() (*sql.DB, error) {
	dataSourceString := "postgres://" + dbs.username + ":" + dbs.password + "@" + dbs.host + ":5432/" + dbs.database + "?sslmode=disable"
	conn, err := sql.Open("postgres", dataSourceString)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
