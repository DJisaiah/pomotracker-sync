package db

import (
	"context"
	"errors"
)

var (
	ErrUserAlreadyExists = errors.New("User already exists")
	ErrInvalidPassword = errors.New("#TODO")
	ErrInvalidUsername = errors.New("#TODO")
)

func (q *Queries) setupUsers()

func (q *Queries) exists(u string) error

func invalidPassword(p string) bool

func invalidUsername(u string) bool

func (q *Queries) AddUser(u string, p string, pH []byte, s bool, lH bool) error {
	if q.exists(u) != nil {
		return ErrUserAlreadyExists
	} else if invalidPassword(p)  {
		return ErrInvalidPassword
	} else if invalidUsername(u) {
		return ErrInvalidUsername
	}

	qry := `

	`	
	q.pool.Exec(context.Background(), qry)

	return nil
	
}
