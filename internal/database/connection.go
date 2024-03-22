package database

import (
	"database/sql"
	"log/slog"

	env "github.com/devjuniorhanun/SisdeveApiGo/config"
	_ "github.com/lib/pq"
)

// Função NewDBConnection()
// Responsável por iniciar uma conexão com o postgreSQL e retornar um ponteiro dessa conexão
func NewDBConnection() (*sql.DB, error) {
	postgresURI := env.Env.DatabaseURL
	db, err := sql.Open("postgres", postgresURI)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	slog.Info("database connected", slog.String("package", "database"))

	return db, nil
}
