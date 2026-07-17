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

func (sa serverActions) registerUser(ac authConfig) (string, string, error) {
	tkn := rand.Text()
	vTil := "somedate"
	slt := make([]byte, 32)
	rand.Read(slt)
	hsh := argon2.IDKey([]byte(ac.Password), slt, 3, 64*1024, 4, 32)
	err := sa.q.AddUser(ac.Username, hsh, ac.Student, ac.LeftHanded)
	if err != nil {
		return "", "", err
	}
	return tkn, vTil, nil
}
