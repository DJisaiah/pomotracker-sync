package db

import (
	"context"
	"errors"
)

var (
	ErrUserAlreadyExists = errors.New("User already exists")
	ErrInvalidPassword = errors.New("#TODO")
	ErrInvalidUsername = errors.New("#TODO")
	ErrFailedToRegister = errors.New("Unable to register")
)

func (q *Queries) setupUsers()

func (q *Queries) exists(u string) error

func invalidPassword(p string) bool

func invalidUsername(u string) bool

func Login() error {
	return nil
}

func (q *Queries) AddUser(u User) error {
	if q.exists(u.LoginDetails.Username) != nil {
		return ErrUserAlreadyExists
	} else if invalidPassword(u.LoginDetails.Password)  {
		return ErrInvalidPassword
	} else if invalidUsername(u.LoginDetails.Username) {
		return ErrInvalidUsername
	}

	qry := `
		INSERT INTO users(username, password, salt, student, leftHanded)
		VALUES ($1, $2, $3, $4, $4)
	`	
	_, err := q.pool.Exec(
		context.Background(),
		qry,
		u.LoginDetails.Username,
		u.LoginCrypt.PasswordHash,
		u.LoginCrypt.Salt,
		u.LoginDetails.Student,
		u.LoginDetails.LeftHanded,
	)
	if err != nil {
		return ErrFailedToRegister 
	}

	err = q.storeToken(u.LoginCrypt.Token, u.LoginCrypt.ValidTil)

	if err != nil {
		return ErrFailedToRegister
	}

	return nil
}

func (q *Queries) storeToken(t string, d string) error {
	qry := `
		INSERT INTO user_sessions()
	`
	return  nil
}
