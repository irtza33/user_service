package postgres

import (
    "database/sql"
    "errors"
    "fmt"
    "log"

    "github.com/irtza33/user_service/internal/domain"
    "github.com/irtza33/user_service/internal/repository"
)

type UserStore interface {
    GetUser(userID int) (string, error)
    CreateUser(name string) (int, error)
    DeleteUser(userID int) error
}

type PostgresUserStore struct {
    db *sql.DB
}

func NewPostgresUserStore(db *sql.DB) *PostgresUserStore {
    return &PostgresUserStore{db: db}
}

func (store *PostgresUserStore) GetUser(userID int) (string, error) {
    var name string
    err := store.db.QueryRow("SELECT name FROM users WHERE user_id = $1", userID).Scan(&name)
    if err != nil {
        if err == sql.ErrNoRows {
            return "", errors.New("user not found")
        }
        return "", fmt.Errorf("error querying user: %v", err)
    }
    return name, nil
}

func (store *PostgresUserStore) CreateUser(name string) (int, error) {
    var userID int
    err := store.db.QueryRow("INSERT INTO users (name) VALUES ($1) RETURNING user_id", name).Scan(&userID)
    if err != nil {
        return 0, fmt.Errorf("error creating user: %v", err)
    }
    return userID, nil
}

func (store *PostgresUserStore) DeleteUser(userID int) error {
    result, err := store.db.Exec("DELETE FROM users WHERE user_id = $1", userID)
    if err != nil {
        return fmt.Errorf("error deleting user: %v", err)
    }
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error checking rows affected: %v", err)
    }
    if rowsAffected == 0 {
        return errors.New("user not found")
    }
    return nil
}