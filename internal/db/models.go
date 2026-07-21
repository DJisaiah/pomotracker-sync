package db


type User struct {
	LoginDetails AuthConfig
	LoginCrypt AuthCrypt
}

type AuthConfig struct {
	Username string
	Password string
	Student bool
	LeftHanded bool
}

type AuthCrypt struct {
	PasswordHash []byte
	Salt []byte
	Token string
	ValidTil string
}
