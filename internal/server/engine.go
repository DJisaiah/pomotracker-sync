package server

import (
	"crypto/rand"
	"golang.org/x/crypto/argon2"
)

func registerUser(ac authConfig) (string, string) {
	t := rand.Text()
	v := "somedate"
	s := make([]byte, 32)
	rand.Read(s)
	h := argon2.IDKey([]byte(ac.Password), s, 3, 64*1024, 4, 32)

	return t, v
}
