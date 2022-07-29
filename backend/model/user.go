package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"` // TODO fix to `json:"-"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createAt"`
}

func (u User) Insert(db *sql.DB) error {
	cmd := `INSERT INTO users (name, password, email) VALUES(
		$1, $2, $3);`

	_, err := db.Exec(cmd, u.Name, u.Password, u.Email)
	if err != nil {
		return err
	}
	return nil
}
