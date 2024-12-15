package database

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq"
)

type PostgresDB struct {
    *sql.DB
}

func NewPostgresDB(dataSourceName string) (*PostgresDB, error) {
    db, err := sql.Open("postgres", dataSourceName)
    if err != nil {
        return nil, err
    }

    if err := db.Ping(); err != nil {
        return nil, err
    }

    return &PostgresDB{db}, nil
}

func (db *PostgresDB) Close() error {
    return db.DB.Close()
}