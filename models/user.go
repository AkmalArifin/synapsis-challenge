package models

import (
	"errors"
	"time"

	"github.com/guregu/null/v5"
	"github.com/synapsis-challenge/db"
	"github.com/synapsis-challenge/utils"
)

type User struct {
	ID        int64       `json:"id"`
	Name      null.String `json:"name"`
	Username  null.String `json:"username"`
	Email     null.String `json:"email"`
	Password  null.String `json:"password"`
	CreatedAt NullTime    `json:"created_at"`
	DeletedAt NullTime    `json:"deleted_at"`
}

func GetAllUsers() ([]User, error) {
	query := `
	SELECT id, name, username, email, password, created_at, deleted_at
	FROM users
	WHERE deleted_at IS NULL
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func GetUserByEmail(email string) (User, error) {
	query := `
	SELECT id, name, username, email, password, created_at, deleted_at
	FROM users
	WHERE email = ? AND deleted_at IS NULL
	`

	var user User
	row := db.DB.QueryRow(query, email)
	err := row.Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.DeletedAt)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (u *User) Save() error {
	query := `
	INSERT INTO users(name, username, email, password, created_at)
	VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password.ValueOrZero())
	if err != nil {
		return err
	}

	u.Password.SetValid(hashedPassword)
	u.CreatedAt.SetValue(time.Now())

	result, err := stmt.Exec(u.Name, u.Username, u.Email, u.Password, u.CreatedAt)
	if err != nil {
		return err
	}

	u.ID, err = result.LastInsertId()

	return err
}

func (u *User) ValidateCredentials() error {
	query := `
	SELECT id, name, username, email, password
	FROM users
	WHERE email = ? AND deleted_at IS NULL
	`

	var retrievedUser User
	row := db.DB.QueryRow(query, u.Email.String)
	err := row.Scan(&retrievedUser.ID, &retrievedUser.Name, &retrievedUser.Username, &retrievedUser.Email, &retrievedUser.Password)
	if err != nil {
		return errors.New("credentials invalid")
	}

	isValid := utils.CompareHashPassword(u.Password.String, retrievedUser.Password.String)
	if !isValid {
		return errors.New("credentials invalid")
	}

	u.ID = retrievedUser.ID
	u.Username.SetValid(retrievedUser.Username.String)

	return nil
}
