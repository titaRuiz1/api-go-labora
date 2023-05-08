package config

import(
	"database/sql"
	_ "github.com/lib/pq"

)

func GetConnection() (*sql.DB, error) {
	database, err := sql.Open("postgres", "postgres://postgres:2015@localhost/labora_project_1?sslmode=disable")
	if err != nil {
		return nil, err
	}

	err = database.Ping()
	if err != nil {
		return nil, err
	}

	return database, nil
}