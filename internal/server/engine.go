package server

import (
	"crypto/rand"
	"golang.org/x/crypto/argon2"
	"github.com/DJisaiah/pomotracker-sync/internal/db"
)

type serverActions struct {
	q *db.Queries
}

func StartServer(q *db.Queries) {
	s := serverActions{q: q}
	start(s)
}

func (sa serverActions) registerUser(ac db.AuthConfig) (string, string, error) {
	tkn := rand.Text()
	vTil := "somedate"
	slt := make([]byte, 32)
	rand.Read(slt)
	pH := argon2.IDKey([]byte(ac.Password), slt, 3, 64*1024, 4, 32)

	usr := db.User{
		LoginDetails: ac,
		LoginCrypt: db.AuthCrypt{
			PasswordHash: pH,
			Salt: slt,
			Token: tkn,
			ValidTil: vTil,
		},
	}
	err := sa.q.AddUser(usr)
	if err != nil {
		return "", "", err
	}
	return tkn, vTil, nil
}
