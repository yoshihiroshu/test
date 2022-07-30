package model

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Password  string    `json:"password"` // TODO fix to `json:"-"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createAt"`
}

func (u *User) SetUUID(id string) {
	u.ID = uuid.MustParse(id)
}

func (u *User) SetCreateAt(date string) {
	createdAt, err := time.Parse(time.RFC3339, date)
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	fmt.Println("setCreatedAt: ", createdAt)
	u.CreatedAt = createdAt
}

func (u User) Insert(db *sql.DB) error {
	cmd := `INSERT INTO users (name, password, email) VALUES(
		$1, $2, $3)RETURNING id;`

	stmt, err := db.Prepare(cmd)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var id string
	err = stmt.QueryRow(u.Name, u.Password, u.Email).Scan(&id)
	if err != nil {
		return err
	}

	u.SetUUID(id)
	return nil
}

func (u User) GetAll(db *sql.DB) ([]User, error) {
	cmd := `SELECT * FROM users LIMIT 10;`

	stmt, err := db.Prepare(cmd)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 10)
	for rows.Next() {
		var id, createdAt string
		var user User
		if err = rows.Scan(&id, &user.Name, &user.Password, &user.Email, &createdAt); err != nil {
			log.Fatalf("failed to scan row: %s", err)
		}

		user.SetUUID(id)
		user.SetCreateAt(createdAt)
		users = append(users, user)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return users, nil
}
